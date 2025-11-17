package units

import "pbnPierre/gowarrior/app"

const MAX_HEALTH = 20

type Player interface {
	toUtf8Char() string
	toANSIChar() float64
	perform_turn()
}

type Warrior struct {
	Score        int
	Attack_power int
	Shoot_power  int
	name         string
	Health       int
	Coordinates  app.Coordinates
}

func NewWarrior(name string, Coordinates app.Coordinates) *Warrior {
	w := Warrior{name: name}
	w.Health = MAX_HEALTH
	w.Score = 0
	w.Attack_power = 2
	w.Shoot_power = 3
	w.Coordinates = Coordinates
	return &w
}

func (w *Warrior) String() string {
	if w.name == "" {
		return "Warrior"
	}

	return w.name
}

func (w Warrior) ToUtf8Char() string {
	return "🤺"
}

func (w Warrior) ToANSIChar() string {
	return "@"
}

func (w *Warrior) FeelEnemy() bool {
	return true
}

func (w *Warrior) Walk() {
	w.Coordinates.X += 1
}

func (w *Warrior) Attack() {
}
