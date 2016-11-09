package gameobj

import (
	"github.com/hajimehoshi/ebiten"
)

type Square struct {
	*BaseShape
	Image *ebiten.Image
}

func NewSquare(base *BaseShape, image *ebiten.Image) *Square {
	var s = &Square{
		BaseShape: base,
		Image:     image,
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
	w, h := s.Image.Size()
	halfHeight := float64(h / 2)
	halfWidth := float64(w / 2)
	return int(s.CenterCoordinate.X - halfHeight),
		int(s.CenterCoordinate.Y - halfWidth),
		int(s.CenterCoordinate.X + halfHeight),
		int(s.CenterCoordinate.Y + halfWidth)
}

func (s *Square) Src(i int) (x0, y0, x1, y1 int) {
	w, h := s.Image.Size()
	return 0, 0, w, h
}
