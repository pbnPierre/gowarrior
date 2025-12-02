package game

import (
	"fmt"
	"pbnPierre/gowarrior/app"
)

type Archer struct {
	health      int
	coordinates app.Coordinates
}

func NewArcher(Coordinates app.Coordinates) *Archer {
	a := Archer{coordinates: Coordinates, health: 7}
	return &a
}

func (a Archer) Name() string {
	return "Archer"
}

func (a Archer) ToChar() string {
	return "🏹"
}

func (a Archer) Coordinates() app.Coordinates {
	return a.coordinates
}

func (a Archer) ShootPower() int {
	return 0
}

func (a *Archer) Attacked(power int) {
	a.health -= power
}

func (a Archer) AttackPower() int {
	return 3
}

func (a Archer) Health() int {
	return a.health
}

func (a Archer) PerformTurn(game *Game) {
	fmt.Printf("%s plays\n", a.Name())
	if a.coordinates.IsCloseTo(game.Player.Warrior.Coordinates) {
		game.AttackAt(a.AttackPower(), a.coordinates)
	}
}
