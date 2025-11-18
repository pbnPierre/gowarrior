package game

import (
	"pbnPierre/gowarrior/app"
)

type Player struct {
	warrior Warrior
}

func (p *Player) PlayTurn() {
	p.warrior.Walk()
}

func NewPlayer(name string) *Player {
	Coordinates := app.Coordinates{X: 0, Y: 0}
	warrior := NewWarrior(name, Coordinates)
	p := Player{warrior: *warrior}
	return &p
}
