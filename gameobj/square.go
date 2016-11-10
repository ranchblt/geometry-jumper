package gameobj

import (
	"github.com/hajimehoshi/ebiten"
)

type Square struct {
	*BaseShape
	image *ebiten.Image
}

func NewSquare(base *BaseShape, image *ebiten.Image) *Square {
	var s = &Square{
		BaseShape: base,
		image:     image,
	}
	return s
}

func (s *Square) Update() {
	// squares dont move vertically, only horizontally.
	s.CenterCoordinate.X = s.CenterCoordinate.X - (s.BaseSpeed * s.SpeedModifier)
}

func (s *Square) Len() int {
	return 1
}

func (s *Square) Dst(i int) (x0, y0, x1, y1 int) {
	w, h := s.image.Size()
	halfHeight := h / 2
	halfWidth := w / 2
	return s.CenterCoordinate.X - halfHeight,
		s.CenterCoordinate.Y - halfWidth,
		s.CenterCoordinate.X + halfHeight,
		s.CenterCoordinate.Y + halfWidth
}

func (s *Square) Src(i int) (x0, y0, x1, y1 int) {
	w, h := s.image.Size()
	return 0, 0, w, h
}

func (s *Square) Image() *ebiten.Image {
	return s.image
}
