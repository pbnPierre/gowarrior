package game

func CreateLevel5() *Tower {
	var units = map[Coordinates]Unit{
		*NewCoordinates(2, 0): NewCaptive(*NewCoordinates(2, 0)),
		*NewCoordinates(3, 0): NewArcher(*NewCoordinates(3, 0)),
		*NewCoordinates(4, 0): NewArcher(*NewCoordinates(4, 0)),
		*NewCoordinates(7, 0): NewCaptive(*NewCoordinates(7, 0)),
	}
	return NewTower(
		"You hear cries for help. Captives must need rescuing.",
		"Use warrior.Feel to see if there is a captive and warrior.Rescue to rescue him. Don't attack captives.",
		"Don't forget to constantly check if you're taking damage. Rest until your health is full if you aren't taking damage.",
		Size{Width: 8, Height: 1},
		*NewCoordinates(6, 0),
		units,
		45,
		123,
	)
}
