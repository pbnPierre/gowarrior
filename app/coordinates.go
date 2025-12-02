package app

import "math"

type Coordinates struct {
	X int
	Y int
}

func NewCoordinates(X int, Y int) *Coordinates {
	c := Coordinates{X: X, Y: Y}
	return &c
}

func (c Coordinates) IsCloseTo(c2 Coordinates) bool {
	return math.Abs(float64(c.X-c2.X+c.Y-c2.Y)) <= 1
}

func (c Coordinates) Equals(c2 Coordinates) bool {
	return c.X == c2.X && c.Y == c2.Y
}
