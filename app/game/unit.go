package game

import (
	"pbnPierre/gowarrior/app"
)

type Unit interface {
	AttackPower() int
	ShootPower() int
	Health() int
	Attacked(power int)
	Coordinates() app.Coordinates
	Name() string
	ToChar() string
	PerformTurn(game *Game)
}
