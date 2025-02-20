package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestRollDie checks if the die rolls a valid number between 1 and 6.
func TestRollDie(t *testing.T) {
	// Create a new assertion object
	a := assert.New(t)

	// Roll the die 100 times and check the results
	for i := 0; i < 100; i++ {
		roll := rollDie()
		// Assert that the roll is within the valid range
		a.True(roll >= 1 && roll <= 6, "Invalid roll: %d", roll)
	}
}

func updateScore(p *Player, points int) int {
	p.score += points
	return p.score
}

// TestUpdateScore verifies that the score is updated correctly.
func TestUpdateScore(t *testing.T) {
	a := assert.New(t)

	// Define test cases for score updates
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
		a.Equal(test.expectedScore, score, "Expected score %d, got %d", test.expectedScore, score)
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
			if handleRoll(1, &turnTotal) {
				t.Error("Expected false for roll of 1")
			}
			if !handleRoll(6, &turnTotal) {
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
	if handleRoll(1, &turnTotal) {
		t.Error("Expected false for roll of 1, indicating the player loses all points for this turn")
	}
	if turnTotal != 0 {
		t.Errorf("Expected turn total to be 0 after rolling 1, got %d", turnTotal)
	}

	// Test for a roll of 2
	if !handleRoll(2, &turnTotal) {
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
			result := simulateGames(test.player1Hold, test.player2Hold, test.numGames)
			// Assert that the total wins match the number of games played
			assert.Equal(t, test.numGames, result.p1Wins+result.p2Wins, "Expected %d total games, got %d", test.numGames, result.p1Wins+result.p2Wins)
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
				assert.Equal(t, test.expectedStart, p1Start, "Expected start %d, got %d", test.expectedStart, p1Start)
				assert.Equal(t, test.expectedEnd, p1End, "Expected end %d, got %d", test.expectedEnd, p1End)
			}
		})
	}
}
