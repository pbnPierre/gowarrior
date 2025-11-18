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
