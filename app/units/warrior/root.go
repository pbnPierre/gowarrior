package warrior

import "fmt"

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
}

func NewWarrior(name string) *Warrior {
	p := Warrior{name: name}
	p.Health = MAX_HEALTH
	p.Score = 0
	p.Attack_power = 2
	p.Shoot_power = 3
	return &p
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

func Perform_turn(w Warrior) {
	fmt.Println("Does Nothing")
}
