package tower

import "pbnPierre/gowarrior/app"

type Size struct {
	Width  int
	Height int
}
type Tower struct {
	Description string
	Tip         string
	Clue        string
	time_bonus  int
	ace_score   int
	Size        Size
	Stairs      app.Coordinates
}

func NewTower(description string, tip string, clue string, time_bonus int, ace_score int, size Size, stairs app.Coordinates) *Tower {
	t := Tower{Description: description, Tip: tip, Clue: clue, time_bonus: time_bonus, ace_score: ace_score, Size: size, Stairs: stairs}
	return &t
}
