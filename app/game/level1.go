package game

import (
	"pbnPierre/gowarrior/app"
)

func CreateLevel1() *Tower {
	var units []Unit
	return NewTower(
		"You see before yourself a long hallway with stairs at the end. There is nothing in the way.",
		"Call p.warrior.Walk() to walk forward in the Player 'PlayTurn' method.",
		"",
		15,
		10,
		Size{Width: 8, Height: 1},
		*app.NewCoordinates(7, 0),
		units,
	)
}
