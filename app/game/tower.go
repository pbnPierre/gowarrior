package game

type Size struct {
	Width  int
	Height int
}

type Tower struct {
	Description string
	Tip         string
	Clue        string
	Size        Size
	Stairs      Coordinates
	Units       map[Coordinates]Unit
}

func NewTower(description string, tip string, clue string, size Size, stairs Coordinates, units map[Coordinates]Unit) *Tower {
	t := Tower{Description: description, Tip: tip, Clue: clue, Size: size, Stairs: stairs, Units: units}
	return &t
}

func (t Tower) hasUnitAt(coordinates Coordinates) bool {
	for _, unit := range t.Units {
		if coordinates.Equals(unit.Coordinates()) {
			return true
		}
	}
	return false
}
