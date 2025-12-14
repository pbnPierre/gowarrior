package game

func CreateLevel4() *Tower {
	var units = map[Coordinates]Unit{
		*NewCoordinates(2, 0): NewThickSludge(*NewCoordinates(2, 0)),
		*NewCoordinates(3, 0): NewArcher(*NewCoordinates(3, 0)),
		*NewCoordinates(5, 0): NewThickSludge(*NewCoordinates(5, 0)),
	}
	return NewTower(
		"You can hear bow strings being stretched.",
		"No new abilities this time, but you must be careful not to rest while taking damage. Save a health instance variable and compare it on each turn to see if you're taking damage.",
		"Set health to your current health at the end of the turn. If this is greater than your current health next turn then you know you're taking damage and shouldn't rest.",
		Size{Width: 7, Height: 1},
		*NewCoordinates(6, 0),
		units,
	)
}
