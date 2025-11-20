package game

import (
	"strings"
	"testing"
)

func TestPrintingLevel1Map(t *testing.T) {
	game := NewGame("José", 1)
	mapDisplay := strings.Trim(game.getMap(), "\n")
	expected := `🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱
🧱🤺🟩🟩🟩🟩🟩🟩📈🧱
🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱`
	if mapDisplay != expected {
		t.Errorf("Level 1 map must display correctly %s does not equals %s", mapDisplay, expected)
	}
	game2 := NewGame("", 2)
	mapDisplay2 := strings.Trim(game2.getMap(), "\n")
	expected2 := `🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱
🧱🤺🟩🟩🟩💩🟩🟩📈🧱
🧱🧱🧱🧱🧱🧱🧱🧱🧱🧱`
	if mapDisplay2 != expected2 {
		t.Errorf("Level 2 map must display correctly %s does not equals %s", mapDisplay2, expected2)
	}
}

func TestPrintingLegend(t *testing.T) {
	game := NewGame("José", 1)
	legend := strings.Trim(game.legend(), "\n")
	expected := `🧱 = Wall
🟩 = Ground
📈 = Stairs
🤺 = José(20 HP)`
	if legend != expected {
		t.Errorf("Level 1 legend must display correctly %s does not equals %s", legend, expected)
	}
	game2 := NewGame("", 2)
	legend2 := strings.Trim(game2.legend(), "\n")
	expected2 := `🧱 = Wall
🟩 = Ground
📈 = Stairs
💩 = Sludge(12 HP)
🤺 = Warrior(20 HP)`
	if legend2 != expected2 {
		t.Errorf("Level 2 legend must display correctly %s does not equals %s", legend2, expected2)
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
		t.Errorf("game 1 must not be considered as same as game 3")
	}

	if game2.isSame(game3) {
		t.Errorf("game 2 must not be considered as same as game 3")
	}

	if !game1.isSame(game1) {
		t.Errorf("game 1 must be considered as same as itself")
	}
}
