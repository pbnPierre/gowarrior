package level1

import (
	"pbnPierre/gowarrior/app"
	"pbnPierre/gowarrior/app/tower"
	"pbnPierre/gowarrior/app/unit"
)

func Create() *tower.Tower {
	var units []unit.Unit
	return tower.NewTower(
		"You see before yourself a long hallway with stairs at the end. There is nothing in the way.",
		"Call p.warrior.Walk() to walk forward in the Player 'PlayTurn' method.",
		"",
		15,
		10,
		tower.Size{Width: 8, Height: 1},
		*app.NewCoordinates(7, 0),
		units,
	)
}
