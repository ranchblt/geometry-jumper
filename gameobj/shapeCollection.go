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
}

func (s *ShapeCollection) Add(g Drawable) {
	s.shapes = append(s.shapes, g)
}

func (s *ShapeCollection) Draw(screen *ebiten.Image) {
	for _, d := range s.shapes {
		screen.DrawImage(d.Image(), &ebiten.DrawImageOptions{
			ImageParts: d,
		})
	}
}
