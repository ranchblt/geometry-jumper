package gameobj

import "github.com/hajimehoshi/ebiten"

type ShapeCollection struct {
	shapes []Drawable
	num    int
}

func NewShapeCollection() *ShapeCollection {
	var s = &ShapeCollection{
		shapes: []Drawable{},
	}
	return s
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
