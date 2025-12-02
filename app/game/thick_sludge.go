package game

import (
	"fmt"
	"pbnPierre/gowarrior/app"
)

type ThickSludge struct {
	health      int
	coordinates app.Coordinates
}

func NewThickSludge(Coordinates app.Coordinates) *ThickSludge {
	s := ThickSludge{coordinates: Coordinates, health: 24}
	return &s
}

func (s ThickSludge) Name() string {
	return "ThickSludge"
}

func (s ThickSludge) ToChar() string {
	return "🌰"
}

func (s ThickSludge) Coordinates() app.Coordinates {
	return s.coordinates
}

func (s ThickSludge) ShootPower() int {
	return 0
}

func (s *ThickSludge) Attacked(power int) {
	s.health -= power
}

func (s ThickSludge) AttackPower() int {
	return 5
}

func (s ThickSludge) Health() int {
	return s.health
}

func (s ThickSludge) PerformTurn(game *Game) {
	fmt.Printf("%s plays\n", s.Name())
	if s.coordinates.IsCloseTo(game.Player.Warrior.Coordinates) {
		game.AttackAt(s.AttackPower(), s.coordinates)
	}

}
