package game

import (
	"fmt"
)

type Sludge struct {
	health      int
	coordinates Coordinates
}

func NewSludge(Coordinates Coordinates) *Sludge {
	s := Sludge{coordinates: Coordinates, health: 12}
	return &s
}

func (s Sludge) Name() string {
	return "Sludge"
}

func (s Sludge) ToChar() string {
	return "💧"
}

func (s Sludge) Coordinates() Coordinates {
	return s.coordinates
}

func (s *Sludge) Attacked(power int) {
	s.health -= power
	fmt.Printf("%s is attacked and loss -%d HP (%d HP)\n", s.Name(), power, s.health)
}

func (s Sludge) AttackPower() int {
	return 3
}

func (s Sludge) Health() int {
	return s.health
}

func (s Sludge) PerformTurn(game *Game) {
	fmt.Printf("%s plays\n", s.Name())
	if s.coordinates.IsCloseTo(game.Player.Warrior.Coordinates) {
		game.AttackAt(game.Player.Warrior.Coordinates, s.AttackPower())
	}

}
