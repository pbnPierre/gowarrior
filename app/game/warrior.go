package game

import (
	"pbnPierre/gowarrior/app"
)

const MAX_HEALTH = 20

type Warrior struct {
	Score        int
	Attack_power int
	Shoot_power  int
	Name         string
	Health       int
	Coordinates  app.Coordinates
}

func NewWarrior(name string, Coordinates app.Coordinates) *Warrior {
	if name == "" {
		name = "Warrior"
	}
	w := Warrior{Name: name, Coordinates: Coordinates}
	w.Health = MAX_HEALTH
	w.Score = 0
	w.Attack_power = 2
	w.Shoot_power = 3
	return &w
}

func (w Warrior) ToChar() string {
	return "🤺"
}

func (w *Warrior) FeelEnemy() bool {
	return true
}

func (w *Warrior) Walk() {
	w.Coordinates.X += 1
}

func (w *Warrior) Attack() {
}
