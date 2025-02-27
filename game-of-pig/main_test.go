package main

import (
	"testing"
)

// TestRollDie checks if the die rolls a valid number between 1 and 6.
func TestRollDie(t *testing.T) {
	for i := 0; i < 100; i++ {
		roll := rollDie()
		if roll < 1 || roll > 6 {
			t.Errorf("Invalid roll: %d", roll)
		}
	}
}

func updateScore(p *Player, points int) int {
	p.score += points
	return p.score
}

// TestUpdateScore verifies that the score is updated correctly.
func TestUpdateScore(t *testing.T) {
	tests := []struct {
		initialScore  int
		pointsToAdd   int
		expectedScore int
	}{
		{0, 2, 2},   // Starting from 0, adding 2 should result in 2
		{10, 3, 13}, // Starting from 10, adding 3 should result in 13
		{20, 6, 26}, // Starting from 20, adding 6 should result in 26
	}

	// Iterate through each test case
	for _, test := range tests {
		// Update the score and assert the expected outcome
		score := updateScore(&Player{score: test.initialScore}, test.pointsToAdd)
		if score != test.expectedScore {
			t.Errorf("Expected score %d, got %d", test.expectedScore, score)
		}
	}
}

func TestGameMechanics(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{"UpdateScore", func(t *testing.T) {
			p := &Player{score: 10}
			if score := updateScore(p, 5); score != 15 {
				t.Errorf("Expected score 15, got %d", score)
			}
		}},
		{"HandleRoll", func(t *testing.T) {
			var turnTotal int
			if handlePigRoll(1, &turnTotal) {
				t.Error("Expected false for roll of 1")
			}
			if !handlePigRoll(6, &turnTotal) {
				t.Error("Expected true for roll of 6")
			}
		}},
	}

	for _, test := range tests {
		t.Run(test.name, test.testFunc)
	}
}

// TestHandleRoll checks the behavior of the handleRoll function.
func TestHandleRoll(t *testing.T) {
	var turnTotal int

	// Test for a roll of 1 (false roll)
	if handlePigRoll(1, &turnTotal) {
		t.Error("Expected false for roll of 1, indicating the player loses all points for this turn")
	}
	if turnTotal != 0 {
		t.Errorf("Expected turn total to be 0 after rolling 1, got %d", turnTotal)
	}

	// Test for a roll of 2
	if !handlePigRoll(2, &turnTotal) {
		t.Error("Expected true for roll of 2, indicating a successful roll")
	}
	if turnTotal != 2 {
		t.Errorf("Expected turn total to be 2 after rolling 2, got %d", turnTotal)
	}
}

// TestGameSimulation verifies the game simulation results for different strategies.
func TestGameSimulation(t *testing.T) {
	tests := []struct {
		name        string
		player1Hold int
		player2Hold int
		numGames    int
	}{
		{"Story1", 10, 15, 10}, // Testing with player 1 holding at 10 and player 2 at 15
		{"Story2", 21, 22, 10}, // Testing with player 1 holding at 21 and player 2 at 22
		{"Story3", 50, 51, 10}, // Testing with player 1 holding at 50 and player 2 at 51
	}

	// Iterate through each test case
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := simulateGames(test.player1Hold, test.player2Hold, test.numGames)
			if result.p1Wins+result.p2Wins != test.numGames {
				t.Errorf("Expected %d total games, got %d", test.numGames, result.p1Wins+result.p2Wins)
			}
		})
	}
}

// TestInputParsing checks the input parsing functionality for various scenarios.
func TestInputParsing(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectError   bool
		expectedStart int
		expectedEnd   int
	}{
		{"NoParameters", []string{}, true, 0, 0},
		{"EmptyParameters", []string{""}, true, 0, 0},
		{"InvalidString", []string{"abc", "def"}, true, 0, 0},
		{"SingleValue", []string{"21", "15"}, false, 21, 21},
		{"Range", []string{"1-100", "1-100"}, false, 1, 100},
		{"InvalidRange", []string{"10-5", "1-100"}, true, 0, 0},
		{"Story1", []string{"10", "15"}, false, 10, 10},
		{"Story2", []string{"21", "1-100"}, false, 21, 21},
		{"Story3", []string{"15-100", "15"}, false, 15, 100},
		{"Story4", []string{"1-100", "1-100"}, false, 1, 100},
	}

	// Iterate through each test case
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p1Start, p1End, _, _, err := parseInput(test.args)
			if test.expectError && err == nil {
				t.Error("Expected error, got nil")
			}
			if !test.expectError {
				if p1Start != test.expectedStart {
					t.Errorf("Expected start %d, got %d", test.expectedStart, p1Start)
				}
				if p1End != test.expectedEnd {
					t.Errorf("Expected end %d, got %d", test.expectedEnd, p1End)
				}
			}
		})
	}
}

// TestHoldingAtInvalidValues checks that holding strategies greater than 100 are not allowed.
func TestHoldingAtInvalidValues(t *testing.T) {
	tests := []struct {
		name        string
		p1Hold      int
		p2Hold      int
		expectError bool
	}{
		{"HoldAt101", 101, 50, true},
		{"HoldAt150", 150, 100, true},
		{"HoldAt0", 0, 50, true},
		{"HoldAtNegative", -10, 50, true},
		{"ValidHold", 50, 50, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := simulateGames(test.p1Hold, test.p2Hold, 10)
			if test.expectError && err == nil {
				t.Error("Expected error, got nil")
			}
		})
	}
}
