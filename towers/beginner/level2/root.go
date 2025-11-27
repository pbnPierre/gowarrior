package level2

import (
	"pbnPierre/gowarrior/app"
	"pbnPierre/gowarrior/app/game/tower"
	"pbnPierre/gowarrior/app/game/unit"
)

func Create() *tower.Tower {
	var units = []unit.Unit{
		unit.NewSludge(*app.NewCoordinates(4, 0)),
	}
	return tower.NewTower(
		"It is too dark to see anything, but you smell sludge nearby.",
		"Use warrior.Feel to see if there is anything in front of you, and warrior.Attack to fight it. Remember, you can only do one action (walk or attack) per turn.",
		"Add an if/else condition using warrior.Feel to decide whether to warrior.Attack or warrior.Walk.",
		15,
		10,
		tower.Size{Width: 8, Height: 1},
		*app.NewCoordinates(7, 0),
		units[:],
	)
}
