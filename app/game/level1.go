package game

func CreateLevel1() *Tower {
	var units map[Coordinates]Unit
	return NewTower(
		"You see before yourself a long hallway with stairs at the end. There is nothing in the way.",
		"Call p.warrior.Walk() to walk forward in the Player 'PlayTurn' method.",
		"",
		Size{Width: 8, Height: 1},
		*NewCoordinates(7, 0),
		units,
	)
}
