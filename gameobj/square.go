package gameobj

type Square struct {
	*BaseShape
}

func NewSquare(track int, centerX float64, centerY float64, baseSpeed float64, speedModifier float64) *Square {
	var s = &Square{
		NewBaseShape(track, centerX, centerY, baseSpeed, speedModifier),
	}
	return s
}

func (s *Square) Update() {
	// squares dont move vertically, only horizontally.
	s.centerX = s.centerX - (s.baseSpeed * s.speedModifier)
}
