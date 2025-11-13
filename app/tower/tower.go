package tower

import "strings"

const WALL = "🧱"
const GROUND = "🟩"
const STAIRS = "📈"

type Coordinates struct {
	X int
	Y int
}
type Size struct {
	Width  int
	Height int
}
type Tower struct {
	description string
	tip         string
	time_bonus  int
	ace_score   int
	size        Size
	stairs      Coordinates
}

func NewTower(description string, tip string, time_bonus int, ace_score int, size Size, stairs Coordinates) *Tower {
	t := Tower{description: description, tip: tip, time_bonus: time_bonus, ace_score: ace_score, size: size, stairs: stairs}
	return &t
}

func (t *Tower) String() string {
	var rows []string
	rows = append(rows, strings.Repeat(WALL, t.size.Width+2))
	for y := 0; y < t.size.Height; y++ {
		var piece []string
		piece = append(piece, WALL)
		for x := 0; x < t.size.Width; x++ {
			if x == t.stairs.X && y == t.stairs.Y {
				piece = append(piece, STAIRS)
			} else {
				piece = append(piece, GROUND)
			}
		}
		piece = append(piece, WALL)
		rows = append(rows, strings.Join(piece, ""))
	}
	rows = append(rows, strings.Repeat(WALL, t.size.Width+2))
	return strings.Join(rows, "\n")
}
