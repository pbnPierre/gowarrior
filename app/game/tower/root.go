package tower

import (
	"pbnPierre/gowarrior/app"
	"pbnPierre/gowarrior/app/game/unit"
)

type Tower struct {
	Description string
	Tip         string
	Clue        string
	TimeBonus   int
	AceScore    int
	Size        Size
	Stairs      app.Coordinates
	Units       []unit.Unit
}

func NewTower(description string, tip string, clue string, time_bonus int, ace_score int, size Size, stairs app.Coordinates, units []unit.Unit) *Tower {
	t := Tower{Description: description, Tip: tip, Clue: clue, TimeBonus: time_bonus, AceScore: ace_score, Size: size, Stairs: stairs, Units: units}
	return &t
}
