package gameobj

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type ShapeCollection struct {
	shapeImageMap   map[int][]*ebiten.Image
	hitboxImageMap  map[int]*ebiten.Image
	shapes          []Drawable
	upperSpeedLimit int
	minimumSpeed    int
	shapeRandom     *rand.Rand
	colorSwap       bool
}

func NewShapeCollection() *ShapeCollection {
	shapeSource := rand.NewSource(time.Now().UnixNano())
	var s = &ShapeCollection{
		shapeImageMap:   ShapeImageMap,
		hitboxImageMap:  HitboxImageMap,
		shapes:          []Drawable{},
		upperSpeedLimit: StartingUpperSpeedLimit,
		minimumSpeed:    MinimumSpeed,
		shapeRandom:     rand.New(shapeSource),
		colorSwap:       false,
	}
	return s
}

func (s *ShapeCollection) IncreaseUpperSpeedLimit() {
	s.upperSpeedLimit++
}

func (s *ShapeCollection) SpawnShape(shapeType int) {
	track := s.shapeRandom.Intn(len(TrackMappings)) + 1
	speed := s.shapeRandom.Intn(s.upperSpeedLimit) + s.minimumSpeed
	image := s.shapeImageMap[shapeType][0]
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

func (s *ShapeCollection) SpawnRandomShape() {
	// this should hopefully give us 1 to len(ShapeTypes)
	shapeType := s.shapeRandom.Intn(len(ShapeTypes)) + 1
	s.SpawnShape(shapeType)
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
