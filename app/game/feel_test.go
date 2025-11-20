package game

import (
	"pbnPierre/gowarrior/app"
	"testing"
)

func TestFeelMonster(t *testing.T) {
	game := NewGame("José", 2)
	feel50 := FeelCoordinates(*game, *app.NewCoordinates(5, 0))
	feel10 := FeelCoordinates(*game, *app.NewCoordinates(1, 0))

	if !feel50.monster {
		t.Errorf("game level 2 must feel a monster on 5.0")
	}
	if feel10.monster {
		t.Errorf("game level 2 must not feel a monster on 1.0")
	}
}

func TestFeelWarriorMonster(t *testing.T) {
	game := NewGame("José", 2)
	feel50 := FeelCoordinates(*game, *app.NewCoordinates(5, 0))
	feel10 := FeelCoordinates(*game, *app.NewCoordinates(1, 0))

	if feel50.warrior {
		t.Errorf("game level 2 must noy feel a warrior on 5.0")
	}
	if !feel10.warrior {
		t.Errorf("game level 2 must feel a warrior on 1.0")
	}
}
