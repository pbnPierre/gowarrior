package game

import (
	"fmt"
	"strconv"
	"strings"

	"pbnPierre/gowarrior/towers/beginner/level1"

	"pbnPierre/gowarrior/app/tower"
)

const MAX_LOOP = 10

const WALL = "🧱"
const GROUND = "🟩"
const STAIRS = "📈"

type Game struct {
	tower  tower.Tower
	player Player
}

func NewGame(name string, level int) *Game {
	tower := level1.Create()
	player := NewPlayer(name)
	g := Game{tower: *tower, player: *player}
	return &g
}

func (g *Game) Run() {
	fmt.Println(g.tower.Description)
	fmt.Println(g.Map())
	fmt.Println(g.Legend())
	fmt.Println(g.tower.Tip)
	i := 0
	for ok := true; ok; ok = !g.HasWon() {
		if i > MAX_LOOP {
			panic("Too much turn boom")
		}
		g.player.PlayTurn()
		i++
		fmt.Println(g.Map())
	}
	fmt.Println("YOU WIN !")
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
			} else if x == g.player.warrior.Coordinates.X && y == g.player.warrior.Coordinates.Y {
				piece = append(piece, g.player.warrior.ToUtf8Char())
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

	rows = append(rows, g.player.warrior.ToUtf8Char()+" = Warrior("+strconv.Itoa(g.player.warrior.Health)+" HP)")

	return strings.Join(rows, "\n")
}

func (g *Game) HasWon() bool {
	return g.tower.Stairs == g.player.warrior.Coordinates
}
