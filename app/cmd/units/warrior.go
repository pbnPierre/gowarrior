package units

import "fmt"

const MAX_HEALTH = 20

type Player interface {
	toUtf8Char() string
	toANSIChar() float64
	perform_turn()
}

type warrior struct {
	score        int
	attack_power int
	shoot_power  int
	name         string
}

func NewWarrior(name string) *warrior {
	p := warrior{name: name}
	p.score = 0
	p.attack_power = 2
	p.shoot_power = 3
	return &p
}

func (w *warrior) String() string {
	if w.name == "" {
		return "Warrior"
	}

	return w.name
}

func (w warrior) ToUtf8Char() string {
	return "🤺"
}

func (w warrior) ToANSIChar() string {
	return "@"
}

func Perform_turn(w warrior) {
	fmt.Println("Does Nothing")
}
