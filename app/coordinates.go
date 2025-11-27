package app

type Coordinates struct {
	X int
	Y int
}

func NewCoordinates(X int, Y int) *Coordinates {
	c := Coordinates{X: X, Y: Y}
	return &c
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (c Coordinates) IsCloseTo(c2 Coordinates) bool {
	diff := c.X - c2.X + c.Y - c2.Y
	return abs(diff) <= 1
}
