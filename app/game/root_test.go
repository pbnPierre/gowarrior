package game

import (
	"pbnPierre/gowarrior/app"
	"strings"
	"testing"
)

func TestPrintingLevel1Map(t *testing.T) {
	tower := NewGame("José", 1)
	mapDisplay := strings.Trim(tower.getMap(), "\n")
	expected := `🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱
🧱🤺🟩🟩🟩🟩🟩🟩📈🧱
🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱`
	if mapDisplay != expected {
		t.Errorf("Level 1 map must display correctly %s does not equals %s", mapDisplay, expected)
	}
}

func TestPrintingLegend(t *testing.T) {
	tower := NewGame("José", 1)
	legend := strings.Trim(tower.legend(), "\n")
	expected := `🧱 = Wall
🟩 = Ground
📈 = Stairs
🤺 = Warrior(20 HP)`
	if legend != expected {
		t.Errorf("Level 1 legend must display correctly %s does not equals %s", legend, expected)
	}
}
func TestSameCompare(t *testing.T) {
	game1 := NewGame("José", 1)
	game2 := NewGame("José", 1)
	game3 := NewGame("Georges", 1)

	if !game1.isSame(game2) {
		t.Errorf("game 1 must be considered as same as game 2")
	}

	if game1.isSame(game3) {
		t.Errorf("game 1 must be considered as same as game 3")
	}

	if !game2.isSame(game3) {
		t.Errorf("game 2 must not be considered as same as game 3")
	}

	if !game1.isSame(game1) {
		t.Errorf("game 1 must not be considered as same as itself")
	}
}

func TestFeelMonster(t *testing.T) {
	game := NewGame("José", 2)
	feel50 := game.Feel(*app.NewCoordinates(5, 0))
	feel10 := game.Feel(*app.NewCoordinates(1, 0))

	if !feel50.monster {
		t.Errorf("game level 2 must feel a monster on 5.0")
	}
	if feel10.monster {
		t.Errorf("game level 2 must not feel a monster on 1.0")
	}
}

func TestFeelWarriorMonster(t *testing.T) {
	game := NewGame("José", 2)
	feel50 := game.Feel(*app.NewCoordinates(5, 0))
	feel10 := game.Feel(*app.NewCoordinates(1, 0))

	if feel50.warrior {
		t.Errorf("game level 2 must noy feel a warrior on 5.0")
	}
	if !feel10.warrior {
		t.Errorf("game level 2 must feel a warrior on 1.0")
	}
}
