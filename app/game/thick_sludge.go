package game

import (
	"fmt"
)

type ThickSludge struct {
	health      int
	coordinates Coordinates
}

func NewThickSludge(Coordinates Coordinates) *ThickSludge {
	s := ThickSludge{coordinates: Coordinates, health: 24}
	return &s
}

func (s ThickSludge) Name() string {
	return "ThickSludge"
}

func (s ThickSludge) ToChar() string {
	return "🌰"
}

func (s ThickSludge) Coordinates() Coordinates {
	return s.coordinates
}

func (s *ThickSludge) Attacked(power int) {
	s.health -= power
	fmt.Printf("%s is attacked and loss -%d HP (%d HP)\n", s.Name(), power, s.health)
}

func (s ThickSludge) AttackPower() int {
	return 3
}

func (s ThickSludge) Health() int {
	return s.health
}

func (s ThickSludge) PerformTurn(game *Game) {
	fmt.Printf("%s plays\n", s.Name())
	if s.coordinates.IsCloseTo(game.Player.Warrior.Coordinates) {
		game.AttackAt(game.Player.Warrior.Coordinates, s.AttackPower())
	}

}
