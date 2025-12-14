package game

func CreateLevel3() *Tower {
	var units = map[Coordinates]Unit{
		*NewCoordinates(2, 0): NewSludge(*NewCoordinates(2, 0)),
		*NewCoordinates(4, 0): NewSludge(*NewCoordinates(4, 0)),
		*NewCoordinates(5, 0): NewSludge(*NewCoordinates(5, 0)),
		*NewCoordinates(7, 0): NewSludge(*NewCoordinates(7, 0)),
	}
	return NewTower(
		"The air feels thicker than before. There must be a horde of sludge.",
		"Be careful not to die! Use warrior.Health to keep an eye on your health, and warrior.Heal() to earn 10% of max health back.",
		"When there is no enemy ahead of you call warrior.rest! until health is full before walking forward.",
		Size{Width: 9, Height: 1},
		*NewCoordinates(8, 0),
		units,
	)
}
