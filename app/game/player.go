package game

import (
	"pbnPierre/gowarrior/app"
	"pbnPierre/gowarrior/app/unit"
)

type Player struct {
	warrior unit.Warrior
}

func (p *Player) PlayTurn() {
	p.warrior.Walk()
}

func NewPlayer(name string) *Player {
	Coordinates := app.Coordinates{X: 0, Y: 0}
	warrior := unit.NewWarrior(name, Coordinates)
	p := Player{warrior: *warrior}
	return &p
}
