package gameobj

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type ShapeCollection struct {
	shapeImageMap map[int]*ebiten.Image
	shapes        []Drawable
	speedModifier int
	shapeRandom   *rand.Rand
}

func NewShapeCollection(shapeImageMap map[int]*ebiten.Image) *ShapeCollection {
	shapeSource := rand.NewSource(time.Now().UnixNano())
	var s = &ShapeCollection{
		shapeImageMap: shapeImageMap,
		shapes:        []Drawable{},
		speedModifier: DefaultSpeedModifier,
		shapeRandom:   rand.New(shapeSource),
	}
	return s
}

func (s *ShapeCollection) IncreaseSpeedModifier() {
	s.speedModifier++
}

func (s *ShapeCollection) SpawnShape(shapeType int) {
	track := s.shapeRandom.Intn(len(TrackMappings)) + 1
	speed := s.shapeRandom.Intn(s.speedModifier) + 1
	image := s.shapeImageMap[shapeType]

	switch shapeType {
	case TriangleType:
		triangle := NewTriangle(NewBaseShape(track, RightSide, speed, image))
		s.Add(triangle)
	case CircleType:
		circle := NewCircle(NewBaseShape(track, RightSide, speed, image))
		s.Add(circle)
	case SquareType:
		square := NewSquare(NewBaseShape(track, RightSide, speed, image))
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
