package game

import (
	"strconv"
	"strings"

	"pbnPierre/gowarrior/towers/beginner/level1"

	"pbnPierre/gowarrior/app/tower"
	"pbnPierre/gowarrior/app/units/warrior"
)

const WALL = "🧱"
const GROUND = "🟩"
const STAIRS = "📈"

type Game struct {
	tower         tower.Tower
	warrior       warrior.Warrior
	warrior_place Coordinates
}

func NewGame(name string, level int) *Game {
	tower := level1.Create()
	warrior := warrior.NewWarrior(name)
	warrior_place := Coordinates{X: 0, Y: 0}
	g := Game{tower: *tower, warrior: *warrior, warrior_place: warrior_place}
	return &g
}

func (g *Game) Map() string {
	var rows []string
	rows = append(rows, strings.Repeat(WALL, g.tower.Size.Width+2))
	for y := 0; y < g.tower.Size.Height; y++ {
		var piece []string
		piece = append(piece, WALL)
		for x := 0; x < g.tower.Size.Width; x++ {
			if x == g.tower.Stairs.X && y == g.tower.Stairs.Y {
				piece = append(piece, STAIRS)
			} else if x == g.warrior_place.X && y == g.warrior_place.Y {
				piece = append(piece, g.warrior.ToUtf8Char())
			} else {
				piece = append(piece, GROUND)
			}
		}
		piece = append(piece, WALL)
		rows = append(rows, strings.Join(piece, ""))
	}
	rows = append(rows, strings.Repeat(WALL, g.tower.Size.Width+2))
	return strings.Join(rows, "\n")
}

func (g *Game) Legend() string {
	var rows []string

	rows = append(rows, WALL+" = Wall")
	rows = append(rows, GROUND+" = Ground")
	rows = append(rows, STAIRS+" = Stairs")

	rows = append(rows, g.warrior.ToUtf8Char()+" = Warrior("+strconv.Itoa(g.warrior.Health)+" HP)")

	return strings.Join(rows, "\n")
}
