package game

import (
	"pbnPierre/gowarrior/app"
)

type Player struct {
	Warrior Warrior
}

func (p *Player) PlayTurn(game Game) {
	p.Warrior.Walk(game)
}

func NewPlayer(name string) *Player {
	Coordinates := app.Coordinates{X: 0, Y: 0}
	warrior := NewWarrior(name, Coordinates)
	p := Player{Warrior: *warrior}
	return &p
}
