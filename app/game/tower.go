package game

import (
	"pbnPierre/gowarrior/app"
)

type Size struct {
	Width  int
	Height int
}

type Tower struct {
	Description string
	Tip         string
	Clue        string
	Size        Size
	Stairs      app.Coordinates
	Units       map[app.Coordinates]Unit
}

func NewTower(description string, tip string, clue string, size Size, stairs app.Coordinates, units map[app.Coordinates]Unit) *Tower {
	t := Tower{Description: description, Tip: tip, Clue: clue, Size: size, Stairs: stairs, Units: units}
	return &t
}
