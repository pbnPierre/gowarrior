package game

import (
	"fmt"
)

type Archer struct {
	health      int
	coordinates Coordinates
}

func NewArcher(Coordinates Coordinates) *Archer {
	a := Archer{coordinates: Coordinates, health: 7}
	return &a
}

func (a Archer) Name() string {
	return "Archer"
}

func (a Archer) ToChar() string {
	return "🏹"
}

func (a Archer) Coordinates() Coordinates {
	return a.coordinates
}

func (a *Archer) Attacked(power int) {
	a.health -= power
	fmt.Printf("%s is attacked and loss -%d HP (%d HP)\n", a.Name(), power, a.health)
}

func (a Archer) AttackPower() int {
	return 3
}

func (a Archer) Health() int {
	return a.health
}

func (a Archer) PerformTurn(game *Game) {
	fmt.Printf("%s plays\n", a.Name())
	if a.coordinates.CanSeeAt(game.Tower, game.Player.Warrior.Coordinates) && !a.coordinates.IsCloseTo(game.Player.Warrior.Coordinates) {
		game.AttackAt(game.Player.Warrior.Coordinates, a.AttackPower())
	}
}
