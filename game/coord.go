package game

type coord struct {
	x int
	y int
}

func (c coord) X() int {
	return c.x
}

func (c coord) Y() int {
	return c.y
}
