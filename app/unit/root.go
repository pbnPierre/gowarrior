package unit

import "pbnPierre/gowarrior/app"

type Unit interface {
	AttackPower() int
	ShootPower() int
	Health() int
	Coordinates() *app.Coordinates
	Name() string
	ToChar() string
	PerformTurn()
}
