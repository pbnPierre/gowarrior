package game

import (
	"fmt"
)

type Captive struct {
	health      int
	coordinates Coordinates
}

func NewCaptive(Coordinates Coordinates) *Captive {
	c := Captive{coordinates: Coordinates, health: 7}
	return &c
}

func (c Captive) Name() string {
	return "Captive"
}

func (c Captive) ToChar() string {
	return "皿"
}

func (c Captive) Coordinates() Coordinates {
	return c.coordinates
}

func (c *Captive) Attacked(power int) {
	c.health -= power
	fmt.Printf("%s is attacked and loss -%d HP (%d HP)\n", c.Name(), power, c.health)
}

func (c Captive) AttackPower() int {
	return 0
}

func (a Captive) Health() int {
	return a.health
}

func (s Captive) IsCaptive() bool {
	return true
}

func (s Captive) IsFoe() bool {
	return false
}

func (a Captive) PerformTurn(game *Game) {
	fmt.Printf("%s is captive and does nothing\n", a.Name())
}
