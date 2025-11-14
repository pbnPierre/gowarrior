package tower

type Coordinates struct {
	X int
	Y int
}

type Size struct {
	Width  int
	Height int
}
type Tower struct {
	Description string
	Tip         string
	time_bonus  int
	ace_score   int
	Size        Size
	Stairs      Coordinates
}

func NewTower(description string, tip string, time_bonus int, ace_score int, size Size, stairs Coordinates) *Tower {
	t := Tower{Description: description, Tip: tip, time_bonus: time_bonus, ace_score: ace_score, Size: size, Stairs: stairs}
	return &t
}
