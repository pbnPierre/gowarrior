package game

import (
	"fmt"
	"math"
)

const MAX_HEALTH = 20

type Warrior struct {
	attackPower     int
	shootPower      int
	Name            string
	Health          int
	Coordinates     Coordinates
	turnActionCount int
}

func NewWarrior(name string, Coordinates Coordinates) *Warrior {
	if name == "" {
		name = "Warrior"
	}
	w := Warrior{Name: name, Coordinates: Coordinates}
	w.Health = MAX_HEALTH
	w.attackPower = 5
	w.shootPower = 3
	w.turnActionCount = 0
	return &w
}

func (w Warrior) ToChar() string {
	return "🤺"
}

func (w *Warrior) StartTurn() {
	w.turnActionCount = 0
}

func (w *Warrior) Heal() {
	if w.turnActionCount > 0 {
		panic("Warrior cannot make two actions in the same turn")
	}
	healPower := int(MAX_HEALTH * 0.1)
	w.Health = int(math.Min(float64(w.Health+healPower), MAX_HEALTH))
	fmt.Printf("%s heals himself of %d HP (%d HP)\n", w.Name, healPower, w.Health)
	w.turnActionCount++
}
func (w *Warrior) Attacked(power int) {
	w.Health = int(math.Max(float64(w.Health-power), 0))
	fmt.Printf("%s is attacked and loss -%d HP(%d HP)\n", w.Name, power, w.Health)
}

func (w *Warrior) Feel(game Game) Feel {
	return *FeelCoordinates(game, w.Coordinates)
}

func (w *Warrior) Walk(game Game) {
	if w.turnActionCount > 0 {
		panic("Warrior cannot make two actions in the same turn")
	}
	feel := game.Player.Warrior.Feel(game)
	if feel.monster {
		panic(fmt.Sprintf("%s walks right into a monster\n", w.Name))
	}
	w.Coordinates = *NewCoordinates(w.Coordinates.X+1, w.Coordinates.Y)
	w.turnActionCount++
	fmt.Printf("%s walks\n", w.Name)
}

func (w *Warrior) Attack(game *Game) {
	if w.turnActionCount > 0 {
		panic("Warrior cannot make two actions in the same turn")
	}
	game.AttackAt(*NewCoordinates(w.Coordinates.X+1, w.Coordinates.Y), w.attackPower)
	w.turnActionCount++
}
