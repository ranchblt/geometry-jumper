package gameobj

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type ShapeCollection struct {
	shapeImageMap     map[int][]*ebiten.Image
	hitboxImageMap    map[int]*ebiten.Image
	shapes            []Drawable
	upperSpeedLimit   int
	minimumSpeed      int
	shapeRandom       *rand.Rand
	allowColorSwap    bool
	patternCollection *PatternCollection
	currentPattern    *Pattern
	currentSpawn      int
	// mike I don't know how to get split seconds, pls halp?
	spawnTimer     int64
	previousUpdate int64
}

func NewShapeCollection(patternCollection *PatternCollection) *ShapeCollection {
	shapeSource := rand.NewSource(time.Now().UnixNano())
	var s = &ShapeCollection{
		shapeImageMap:     ShapeImageMap,
		hitboxImageMap:    HitboxImageMap,
		shapes:            []Drawable{},
		upperSpeedLimit:   StartingUpperSpeedLimit,
		minimumSpeed:      MinimumSpeed,
		shapeRandom:       rand.New(shapeSource),
		allowColorSwap:    false,
		patternCollection: patternCollection,
		// TODO change me to choose at random
		currentPattern: patternCollection.Patterns[LowDifficulty][0],
		spawnTimer:     0,
		previousUpdate: time.Now().Unix(),
	}
	return s
}

func (s *ShapeCollection) shapeFromSpawn(spawn *Spawn) {
	shapeType := spawn.ShapeType
	track := spawn.Track
	speed := spawn.Speed
	shapeImages := s.shapeImageMap[shapeType]
	var image *ebiten.Image
	if s.allowColorSwap {
		imageIndex := s.shapeRandom.Intn(len(shapeImages))
		image = shapeImages[imageIndex]
	} else {
		image = shapeImages[0]
	}
	hitboxImage := s.hitboxImageMap[shapeType]

	switch shapeType {
	case TriangleType:
		triangle := NewTriangle(NewBaseShape(track, RightSide, speed, image, hitboxImage))
		s.Add(triangle)
	case CircleType:
		circle := NewCircle(NewBaseShape(track, RightSide, speed, image, hitboxImage))
		s.Add(circle)
	case SquareType:
		square := NewSquare(NewBaseShape(track, RightSide, speed, image, hitboxImage))
		s.Add(square)
	}
}

func (s *ShapeCollection) UnlockColorSwap() {
	s.allowColorSwap = true
}

func (s *ShapeCollection) Update() {
	// this should all be in its own method
	currentUpdate := time.Now().Unix()
	s.spawnTimer += currentUpdate - s.previousUpdate
	s.previousUpdate = currentUpdate

	if s.currentPattern.SpawnReady(s.currentSpawn, s.spawnTimer) {
		spawn := s.currentPattern.GetCurrentSpawn(s.currentSpawn)
		s.shapeFromSpawn(spawn)
		s.currentSpawn++
		s.spawnTimer = 0
	}

	if s.currentPattern.OnLastSpawn(s.currentSpawn) {
		fmt.Println("we are done with this pattern")
		// todo: real randomizing for the spawns
		s.currentPattern = s.patternCollection.Patterns[LowDifficulty][0]
		s.currentSpawn = 0
	}
	// end this should all be in its own method

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
