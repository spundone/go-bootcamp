package main

import (
	// "fmt"
	// "os"
	"testing"
)

func TestRollDie(t *testing.T) {
	//Test if value rolled by the die is between 1 and 6
	for i := 0; i < 100; i++ {
		roll := rollDie()
		if roll < 1 || roll > 6 {
			t.Errorf("Invalid roll: %d", roll)
		}
	}
}

func TestCheckPig(t *testing.T) {
	tests := []struct {
		roll int
		pig  bool
	}{
		{1, true}, //dunno if more or less resource intensive compared to running a loop
		{2, false},
		{3, false},
		{4, false},
		{5, false},
		{6, false},
	}
	for _, test := range tests {
		if checkPig(test.roll) != test.pig {
			t.Errorf("checkPig(%d) = %v, want %v", test.roll, !test.pig, test.pig)
		}
	}
}

func TestUpdateScore(t *testing.T) {
	tests := []struct {
		currentScore  int
		roll          int
		expectedScore int
	}{
		{0, 2, 2}, //
		{10, 3, 13},
		{20, 6, 26},
	}
	for _, test := range tests {
		score := updateScore(&Player{score: test.currentScore}, test.roll)
		if score != test.expectedScore {
			t.Errorf("updateScore(%d, %d) = %d, want %d", test.currentScore, test.roll, score, test.expectedScore)
		}
	}
}

func TestGameOfPig(t *testing.T) {
	// Test that game ends when a player reaches 100 points

	// Test that the game switches between players

	// Test that game handles pig rolls correctly
}
