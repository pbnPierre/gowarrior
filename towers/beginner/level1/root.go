package level1

import (
	"pbnPierre/gowarrior/app/tower"
)

func Create() *tower.Tower {
	return tower.NewTower(
		"You see before yourself a long hallway with stairs at the end. There is nothing in the way.",
		"Call warrior.walk! to walk forward in the Player 'play_turn' method.",
		15,
		10,
		tower.Size{Width: 8, Height: 1},
		tower.Coordinates{X: 7, Y: 0},
	)
}
