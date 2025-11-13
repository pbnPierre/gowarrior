package units

import "fmt"

const MAX_HEALTH = 20

type player interface {
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

func newWarrior(name string) *warrior {
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

func (w warrior) toUtf8Char() string {
	return "🤺"
}

func (w warrior) toANSIChar() string {
	return "@"
}

func perform_turn(w warrior) {
	fmt.Println("Does Nothing")
}
