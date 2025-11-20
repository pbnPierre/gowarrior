package unit

import (
	"pbnPierre/gowarrior/app"
)

type Sludge struct {
	health      int
	coordinates app.Coordinates
}

func NewSludge(Coordinates app.Coordinates) *Sludge {
	s := Sludge{coordinates: Coordinates, health: 12}
	return &s
}

func (s Sludge) Name() string {
	return "Sludge"
}

func (s Sludge) ToChar() string {
	return "💩"
}

func (s Sludge) Coordinates() *app.Coordinates {
	return &s.coordinates
}

func (s Sludge) ShootPower() int {
	return 0
}

func (s Sludge) AttackPower() int {
	return 3
}

func (s Sludge) Health() int {
	return s.health
}

func (s Sludge) PerformTurn() {

}
