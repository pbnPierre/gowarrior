package tower

import (
	"pbnPierre/gowarrior/app"
	"pbnPierre/gowarrior/app/unit"
)

type Tower struct {
	Description string
	Tip         string
	Clue        string
	Time_bonus  int
	Ace_score   int
	Size        Size
	Stairs      app.Coordinates
	Units       []unit.Unit
}

func NewTower(description string, tip string, clue string, time_bonus int, ace_score int, size Size, stairs app.Coordinates, units []unit.Unit) *Tower {
	t := Tower{Description: description, Tip: tip, Clue: clue, Time_bonus: time_bonus, Ace_score: ace_score, Size: size, Stairs: stairs, Units: units}
	return &t
}
