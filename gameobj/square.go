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
	s.centerX = s.centerX - (s.baseSpeed * s.speedModifier)
}
