package game

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/uber-go/zap"
)

type ShapeCollection struct {
	shapeImageMap        map[int][]*ebiten.Image
	shapes               []Drawable
	upperSpeedLimit      int
	minimumSpeed         int
	shapeRandom          *rand.Rand
	allowColorSwap       bool
	patternCollection    *PatternCollection
	currentPattern       *Pattern
	duringPattern        bool
	unlockedDifficulties []string
	// used to lock easier difficulties
	difficultyOffset        int
	Stop                    bool
	difficultyChangesQueued bool
}

func NewShapeCollection() *ShapeCollection {
	shapeSource := rand.NewSource(time.Now().UnixNano())
	var s = &ShapeCollection{
		shapeImageMap:        ShapeImageMap,
		shapes:               []Drawable{},
		upperSpeedLimit:      StartingUpperSpeedLimit,
		minimumSpeed:         MinimumSpeed,
		shapeRandom:          rand.New(shapeSource),
		allowColorSwap:       false,
		patternCollection:    GamePatternCollection,
		unlockedDifficulties: []string{DifficultyTypes[0]},
		difficultyOffset:     0,
	}
	return s
}

func (s *ShapeCollection) assignPatternOnDelay(delayMillis int64) {
	// need to assign this prior to the timer finishing, otherwise we'll stack up patterns
	// and cause a hilariously unavoidable line of death.
	s.duringPattern = true
	timer := time.NewTimer(time.Millisecond * time.Duration(delayMillis))
	<-timer.C
	s.assignPattern()
}

// assigns a pattern randomly using the currently unlocked difficulties
func (s *ShapeCollection) assignPattern() {
	difficultyIndex := s.shapeRandom.Intn(len(s.unlockedDifficulties)-s.difficultyOffset) + s.difficultyOffset
	difficulty := DifficultyTypes[difficultyIndex]

	logger.Debug("Difficulty for pattern",
		zap.String("difficulty", difficulty),
	)

	patterns := s.patternCollection.Patterns[difficulty]
	patternIndex := s.shapeRandom.Intn(len(patterns))
	s.currentPattern = patterns[patternIndex]

	spawnGroups := s.currentPattern.SpawnGroups
	for index, spawnGroup := range spawnGroups {
		go s.spawnOnDelay(spawnGroup, index == len(spawnGroups)-1)
	}
}

func (s *ShapeCollection) spawnOnDelay(spawnGroup *SpawnGroup, lastSpawn bool) {
	logger.Debug("spawnOnDelay",
		zap.Int("spawn time miliseconds", spawnGroup.SpawnTimeMillis),
	)

	timer := time.NewTimer(time.Millisecond * time.Duration(spawnGroup.SpawnTimeMillis))
	<-timer.C
	// So shapes down't spawn after
	if s.Stop {
		return
	}
	logger.Debug("Done Waiting")
	for _, spawn := range spawnGroup.Spawns {
		s.shapeFromSpawn(spawn)
	}
	// if it's the last spawn group, we're not on the pattern anymore
	s.duringPattern = !lastSpawn
}

func (s *ShapeCollection) shapeFromSpawn(spawn *Spawn) {
	shapeType := spawn.ShapeType
	track := spawn.Track
	speed := spawn.Speed
	image := s.shapeImageMap[shapeType][0]
	colorMap := s.getImageColorMap(shapeType)

	switch shapeType {
	case TriangleType:
		triangle := NewTriangle(NewBaseShape(track, RightSide, speed, image, colorMap))
		s.Add(triangle)
	case CircleType:
		circle := NewCircle(NewBaseShape(track, RightSide, speed, image, colorMap))
		s.Add(circle)
	case SquareType:
		square := NewSquare(NewBaseShape(track, RightSide, speed, image, colorMap))
		s.Add(square)
	}
}

// gets the image colors for a shape. if color swapping is enabled, this is randomly chosen
// from the shape length
func (s *ShapeCollection) getImageColorMap(shapeType int) ebiten.ColorM {
	var colorMap ebiten.ColorM
	// I dont like reassigning method param values, where possible
	var calculatedShapeType int
	if s.allowColorSwap {
		shapeTypeIndex := s.shapeRandom.Intn(len(ShapeTypes))
		calculatedShapeType = ShapeTypes[shapeTypeIndex]
	} else {
		calculatedShapeType = shapeType
	}
	colorMap = ColorMappings[calculatedShapeType]
	return colorMap
}

func (s *ShapeCollection) unlockColorSwapOnDelay(delaySeconds int64) {
	timer := time.NewTimer(time.Second * time.Duration(delaySeconds))
	<-timer.C
	s.unlockColorSwap()
}

func (s *ShapeCollection) unlockColorSwap() {
	s.allowColorSwap = true
	logger.Debug("Unlocked color swap")
}

func (s *ShapeCollection) unlockDifficultyOnDelay(delaySeconds int64) {
	timer := time.NewTimer(time.Second * time.Duration(delaySeconds))
	<-timer.C
	s.unlockNextDifficulty()
}

func (s *ShapeCollection) unlockNextDifficulty() {
	// difficulties are added sequentially, so just use our length as an index
	// into the difficulty type slice
	if len(s.unlockedDifficulties) < len(DifficultyTypes) {
		nextDifficultyIndex := len(s.unlockedDifficulties)
		nextDifficulty := DifficultyTypes[nextDifficultyIndex]
		s.unlockedDifficulties = append(s.unlockedDifficulties, nextDifficulty)
		logger.Debug("Unlocked",
			zap.String("Difficulty", nextDifficulty),
		)
	} else {
		logger.Debug("No more difficulties to unlock")
	}
}

func (s *ShapeCollection) lockDifficultyOnDelay(delaySeconds int64) {
	timer := time.NewTimer(time.Second * time.Duration(delaySeconds))
	<-timer.C
	s.lockDifficulty()
}

func (s *ShapeCollection) lockDifficulty() {
	// difficulties are added sequentially, so just use our length as an index
	// into the difficulty type slice
	if s.difficultyOffset < len(DifficultyTypes)-1 {
		s.difficultyOffset++
		logger.Debug("Locked",
			zap.String("Difficulty", s.unlockedDifficulties[s.difficultyOffset-1]),
		)
	} else {
		logger.Debug("no more difficulties to lock")
	}
}

func (s *ShapeCollection) Update() {
	for _, d := range s.shapes {
		d.Update()
	}

	// now that we're done updating, let's figure out what shapes
	// expired and remove them
	var unexpiredShapes = []Drawable{}

	for _, d := range s.shapes {
		if !d.IsExpired() {
			unexpiredShapes = append(unexpiredShapes, d)
		}
	}

	if len(s.shapes) == 0 && !s.duringPattern {
		logger.Debug("Assigning new pattern")
		go s.assignPatternOnDelay(PatternDelayMillis)
	}

	if !s.difficultyChangesQueued {
		go s.unlockDifficultyOnDelay(MediumDifficultyUnlockSeconds)
		go s.unlockDifficultyOnDelay(HighDifficultyUnlockSeconds)
		go s.unlockColorSwapOnDelay(ColorSwapUnlockSeconds)
		go s.lockDifficultyOnDelay(LowDifficultyLockSeconds)
		go s.lockDifficultyOnDelay(MediumDifficultyLockSeconds)
		s.difficultyChangesQueued = true
	}

	// boy I hope this doesn't cause a leak somehow
	s.shapes = unexpiredShapes
}

func (s *ShapeCollection) Add(g Drawable) {
	s.shapes = append(s.shapes, g)
}

func (s *ShapeCollection) Draw(screen *ebiten.Image) {
	for _, d := range s.shapes {
		d.Draw(screen)
	}
}
