package game

import (
	"pbnPierre/gowarrior/app"
	"testing"
)

func TestDisplay(t *testing.T) {
	warrior := NewWarrior("", *app.NewCoordinates(0, 0))

	if warrior.Name != "Warrior" {
		t.Errorf("Warrior should have a default value if no data provided")
	}
}

func TestHealing(t *testing.T) {
	warrior := NewWarrior("", *app.NewCoordinates(0, 0))

	warrior.Attacked(3)
	if warrior.Health != (MAX_HEALTH - 3) {
		t.Errorf("An attacked Warrior should loose HP")
	}
	warrior.Heal()
	warrior.Heal()
	warrior.Heal()
	if warrior.Health != MAX_HEALTH {
		t.Errorf("Warrior should not have more than MAX_HEALTH const")
	}
}

func TestAttack(t *testing.T) {
	warrior := NewWarrior("", *app.NewCoordinates(0, 0))

	warrior.Attacked(30)
	if warrior.Health < 0 {
		t.Errorf("Warrior should not have less than 0 health")
	}
}
