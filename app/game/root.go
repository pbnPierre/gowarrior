package game

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"

	"pbnPierre/gowarrior/app"
	"pbnPierre/gowarrior/towers/beginner/level1"
	"pbnPierre/gowarrior/towers/beginner/level2"

	"pbnPierre/gowarrior/app/tower"

	"github.com/huandu/go-clone"
)

const MAX_LOOP = 10

const WALL = "🧱"
const GROUND = "🟩"
const STAIRS = "📈"

type Game struct {
	Tower  tower.Tower
	Player Player
}

func NewGame(name string, level int) *Game {
	towers := map[int]*tower.Tower{
		1: level1.Create(),
		2: level2.Create(),
	}

	player := NewPlayer(name)
	g := Game{Tower: *towers[level], Player: *player}
	return &g
}

func (g *Game) Run() {
	fmt.Println("Hello dear", g.Player.Warrior.Name, "!")
	fmt.Println(g.Tower.Description)
	fmt.Println(g.getMap())
	fmt.Println(g.legend())
	fmt.Println(g.Tower.Tip)
	for ok := true; ok; ok = !g.hasWon() {
		previousState := clone.Clone(g).(*Game)
		g.Player.PlayTurn()
		fmt.Println(g.getMap())
		if g.isSame(previousState) {
			panic("Game state is stuck boom boom")
		}
	}
	fmt.Println("YOU WIN !")
}

func (g Game) getCharForCoordinate(coordinates app.Coordinates) string {
	if coordinates.X == g.Tower.Stairs.X && coordinates.Y == g.Tower.Stairs.Y {
		return STAIRS
	} else if coordinates.X == g.Player.Warrior.Coordinates.X && coordinates.Y == g.Player.Warrior.Coordinates.Y {
		return g.Player.Warrior.ToChar()
	}
	for _, unit := range g.Tower.Units {
		if coordinates.X == unit.Coordinates().X && coordinates.Y == unit.Coordinates().Y {
			return unit.ToChar()
		}
	}
	return GROUND
}

func (g Game) getMap() string {
	var rows []string
	rows = append(rows, strings.Repeat(WALL, g.Tower.Size.Width+2))
	for y := 0; y < g.Tower.Size.Height; y++ {
		var piece []string
		piece = append(piece, WALL)
		for x := 0; x < g.Tower.Size.Width; x++ {
			piece = append(piece, g.getCharForCoordinate(*app.NewCoordinates(x, y)))
		}
		piece = append(piece, WALL)
		rows = append(rows, strings.Join(piece, ""))
	}
	rows = append(rows, strings.Repeat(WALL, g.Tower.Size.Width+2))
	return strings.Join(rows, "\n")
}

func (g Game) legend() string {
	var rows []string

	rows = append(rows, WALL+" = Wall")
	rows = append(rows, GROUND+" = Ground")
	rows = append(rows, STAIRS+" = Stairs")

	for _, unit := range g.Tower.Units {
		rows = append(rows, unit.ToChar()+" = "+unit.Name()+"("+strconv.Itoa(unit.Health())+" HP)")
	}
	rows = append(rows, g.Player.Warrior.ToChar()+" = "+g.Player.Warrior.Name+"("+strconv.Itoa(g.Player.Warrior.Health)+" HP)")

	return strings.Join(rows, "\n")
}

func (g Game) hasWon() bool {
	return g.Tower.Stairs == g.Player.Warrior.Coordinates
}

func (g *Game) hash() []byte {
	hash := md5.New()
	io.WriteString(hash, fmt.Sprint(g.Player.Warrior.Name))
	io.WriteString(hash, fmt.Sprint(g.Player.Warrior.Health))
	io.WriteString(hash, fmt.Sprint(g.Player.Warrior.Coordinates.X))
	io.WriteString(hash, fmt.Sprint(g.Player.Warrior.Coordinates.Y))
	for _, unit := range g.Tower.Units {
		io.WriteString(hash, fmt.Sprint(unit.Health()))
		io.WriteString(hash, fmt.Sprint(unit.Coordinates().X))
		io.WriteString(hash, fmt.Sprint(unit.Coordinates().Y))
	}
	return hash.Sum(nil)
}

func (g *Game) isSame(game *Game) bool {
	return bytes.Equal(g.hash(), game.hash())
}
