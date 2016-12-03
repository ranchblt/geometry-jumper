package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type ShapeCollection struct {
	shapeImageMap          map[int][]*ebiten.Image
	shapes                 []Drawable
	upperSpeedLimit        int
	minimumSpeed           int
	shapeRandom            *rand.Rand
	allowColorSwap         bool
	patternCollection      *PatternCollection
	currentPattern         *Pattern
	currentSpawnTimeTarget time.Time
	unlockedDifficulties   []int
}

func NewShapeCollection(patternCollection *PatternCollection) *ShapeCollection {
	shapeSource := rand.NewSource(time.Now().UnixNano())
	var s = &ShapeCollection{
		shapeImageMap:        ShapeImageMap,
		shapes:               []Drawable{},
		upperSpeedLimit:      StartingUpperSpeedLimit,
		minimumSpeed:         MinimumSpeed,
		shapeRandom:          rand.New(shapeSource),
		allowColorSwap:       false,
		patternCollection:    patternCollection,
		unlockedDifficulties: []int{DifficultyTypes[0]},
	}
	s.assignPattern()
	return s
}

// assigns a pattern randomly using the currently unlocked difficulties
func (s *ShapeCollection) assignPattern() {
	difficultyIndex := s.shapeRandom.Intn(len(s.unlockedDifficulties))
	difficulty := DifficultyTypes[difficultyIndex]

	patterns := s.patternCollection.Patterns[difficulty]
	patternIndex := s.shapeRandom.Intn(len(patterns))
	s.currentPattern = patterns[patternIndex]

	spawnDelay := s.currentPattern.GetCurrentSpawn().SpawnDelayMillis
	s.currentSpawnTimeTarget = time.Now().Add(time.Duration(spawnDelay) * time.Millisecond)

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

func (s *ShapeCollection) UnlockColorSwap() {
	s.allowColorSwap = true
}

func (s *ShapeCollection) UnlockNextDifficulty() {
	// difficulties are added sequentially, so just use our length as an index
	// into the difficulty type slice
	if len(s.unlockedDifficulties) < len(DifficultyTypes) {
		nextDifficultyIndex := len(s.unlockedDifficulties)
		nextDifficulty := DifficultyTypes[nextDifficultyIndex]
		s.unlockedDifficulties = append(s.unlockedDifficulties, nextDifficulty)
	} else {
		fmt.Println("no more difficulties to unlock")
	}
}

func (s *ShapeCollection) updatePattern() {
	now := time.Now()
	if now.After(s.currentSpawnTimeTarget) {
		// spawn the new shape
		spawn := s.currentPattern.GetCurrentSpawn()
		s.shapeFromSpawn(spawn)
		if s.currentPattern.OnLastSpawn() {
			// then if the pattern is finished, reset the old one, get a new one
			s.currentPattern.ResetPattern()
			s.assignPattern()

		} else {
			// otherwise advance the current pattern
			s.currentPattern.AdvancePattern()

		}
		spawnDelay := s.currentPattern.GetCurrentSpawn().SpawnDelayMillis
		// then no matter what happened, we need to move the spawn time target up
		s.currentSpawnTimeTarget = now.Add(time.Duration(spawnDelay) * time.Millisecond)
	}
}

func (s *ShapeCollection) Update() {
	s.updatePattern()

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
