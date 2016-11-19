package gameobj

type Square struct {
	*BaseShape
}

func NewSquare(base *BaseShape) *Square {
	var s = &Square{
		BaseShape: base,
	}
	return s
}

func (s *Square) Update() {
	// squares dont move vertically, only horizontally.
	s.CenterCoordinate.X = s.CenterCoordinate.X - (s.BaseSpeed * s.SpeedModifier)

	if s.crossedLeftEdge() {
		s.expired = true
	}
}
