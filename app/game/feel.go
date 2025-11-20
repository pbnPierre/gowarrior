package game

import "pbnPierre/gowarrior/app"

type Feel struct {
	warrior bool
	monster bool
}

func FeelCoordinates(game Game, c app.Coordinates) *Feel {
	feel := &Feel{monster: false}

	feel.warrior = game.Player.Warrior.Coordinates.IsCloseTo(c)
	for _, unit := range game.Tower.Units {
		feel.monster = feel.monster || unit.Coordinates().IsCloseTo(c)
		if feel.monster {
			break
		}
	}

	return feel
}
