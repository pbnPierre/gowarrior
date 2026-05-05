package game

type Feel struct {
	warrior bool
	captive bool
	monster bool
}

func FeelCoordinates(game Game, c Coordinates) *Feel {
	feel := &Feel{monster: false}

	feel.warrior = game.Player.Warrior.Coordinates.IsCloseTo(c)
	for _, unit := range game.Tower.Units {
		if unit.Coordinates().IsCloseTo(c) {
			if unit.IsCaptive() {
				feel.captive = true
				break
			}
			if unit.IsFoe() {
				feel.monster = true
				break
			}
		}
	}

	return feel
}
