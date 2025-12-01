package game

import (
	"pbnPierre/gowarrior/app"
)

func CreateLevel3() *Tower {
	var units = []Unit{
		NewSludge(*app.NewCoordinates(2, 0)),
		NewSludge(*app.NewCoordinates(4, 0)),
		NewSludge(*app.NewCoordinates(5, 0)),
		NewSludge(*app.NewCoordinates(7, 0)),
	}
	return NewTower(
		"The air feels thicker than before. There must be a horde of sludge.",
		"Be careful not to die! Use warrior.Health to keep an eye on your health, and warrior.Heal() to earn 10% of max health back.",
		"When there is no enemy ahead of you call warrior.rest! until health is full before walking forward.",
		15,
		10,
		Size{Width: 9, Height: 1},
		*app.NewCoordinates(8, 0),
		units[:],
	)
}
