package level2

import (
	"pbnPierre/gowarrior/app"
	"pbnPierre/gowarrior/app/tower"
)

func Create() *tower.Tower {
	return tower.NewTower(
		"It is too dark to see anything, but you smell sludge nearby.",
		"Use warrior.FeelEnemy to see if there is anything in front of you, and warrior.Attack to fight it. Remember, you can only do one action (walk or attack) per turn.",
		"Add an if/else condition using warrior.FeelEnemy to decide whether to warrior.Attack or warrior.Walk.",
		15,
		10,
		tower.Size{Width: 8, Height: 1},
		*app.NewCoordinates(7, 0),
	)
}
