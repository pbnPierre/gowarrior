package game

func CreateLevel2() *Tower {
	var units = map[Coordinates]Unit{
		*NewCoordinates(4, 0): NewSludge(*NewCoordinates(4, 0)),
	}
	return NewTower(
		"It is too dark to see anything, but you smell sludge nearby.",
		"Use warrior.Feel to see if there is anything in front of you, and warrior.Attack to fight it. Remember, you can only do one action (walk or attack) per turn.",
		"Add an if/else condition using warrior.Feel to decide whether to warrior.Attack or warrior.Walk.",
		Size{Width: 8, Height: 1},
		*NewCoordinates(7, 0),
		units,
	)
}
