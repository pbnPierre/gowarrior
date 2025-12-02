package game

import (
	"pbnPierre/gowarrior/app"
)

func CreateLevel4() *Tower {
	var units = map[app.Coordinates]Unit{
		*app.NewCoordinates(2, 0): NewThickSludge(*app.NewCoordinates(2, 0)),
		*app.NewCoordinates(3, 0): NewArcher(*app.NewCoordinates(3, 0)),
		*app.NewCoordinates(5, 0): NewThickSludge(*app.NewCoordinates(5, 0)),
	}
	return NewTower(
		"You can hear bow strings being stretched.",
		"No new abilities this time, but you must be careful not to rest while taking damage. Save a health instance variable and compare it on each turn to see if you're taking damage.",
		"Set health to your current health at the end of the turn. If this is greater than your current health next turn then you know you're taking damage and shouldn't rest.",
		Size{Width: 7, Height: 1},
		*app.NewCoordinates(6, 0),
		units,
	)
}
