package game

import (
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
	tower1 := NewGame("José", 1)
	tower2 := NewGame("José", 1)
	tower3 := NewGame("Georges", 1)

	if tower1.isSame(tower2) {
		t.Errorf("tower1 must be considered as same as tower2")
	}

	if tower1.isSame(tower3) {
		t.Errorf("tower1 must be considered as same as tower3")
	}

	if tower2.isSame(tower3) {
		t.Errorf("tower2 must not be considered as same as tower3")
	}

	if tower1.isSame(tower1) {
		t.Errorf("tower1 must not be considered as same as itself")
	}
}
