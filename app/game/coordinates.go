package game

import (
	"math"
)

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

func (c Coordinates) CanSeeAt(tower Tower, c2 Coordinates) bool {
	if c.X < c2.X {
		for i := c.X + 1; i < c2.X; i++ {
			if tower.hasUnitAt(*NewCoordinates(i, c.Y)) {
				return false
			}
		}
	} else {
		for i := c2.X + 1; i < c.X; i++ {
			if tower.hasUnitAt(*NewCoordinates(i, c.Y)) {
				return false
			}
		}
	}

	return true
}
