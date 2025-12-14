package game

import (
	"testing"
)

func TestCanSeeAt(t *testing.T) {
	game := NewGame(playerJose, level2)
	coordinates10 := *NewCoordinates(1, 0)
	coordinates40 := *NewCoordinates(4, 0)
	coordinates50 := *NewCoordinates(5, 0)
	if coordinates10.CanSeeAt(game.Tower, coordinates50) {
		t.Error("You cannot see through units from left to right")
	}
	if !coordinates10.CanSeeAt(game.Tower, coordinates40) {
		t.Error("You must see in front of you from left to right")
	}
	if !coordinates40.CanSeeAt(game.Tower, coordinates50) {
		t.Error("You must see just next to you from left to right")
	}
	if coordinates50.CanSeeAt(game.Tower, coordinates10) {
		t.Error("You cannot see through units from right to left")
	}
	if !coordinates40.CanSeeAt(game.Tower, coordinates10) {
		t.Error("You must see in front of you from right to left")
	}
	if !coordinates50.CanSeeAt(game.Tower, coordinates40) {
		t.Error("You must see just next to you from right to left")
	}
}
