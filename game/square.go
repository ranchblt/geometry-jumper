package game

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
	s.Center.x = s.Center.x - (s.BaseSpeed)
	s.crossedLeftEdge()
}
