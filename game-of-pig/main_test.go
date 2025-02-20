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
	player := &Player{score: 95}
	roll := 5
	updateScore(player, roll) // Simulate a roll that brings the player to 100
	if player.score != 100 {
		t.Errorf("Expected score to be 100, got %d", player.score)
	}

	// Test that the game switches between players
	player1 := &Player{score: 50}
	player2 := &Player{score: 45}
	currentPlayer := player1
	roll = 3
	updateScore(currentPlayer, roll) // Player 1 rolls
	if currentPlayer.score != 53 {
		t.Errorf("Expected Player 1 score to be 53, got %d", currentPlayer.score)
	}
	currentPlayer = player2 // Switch to Player 2
	roll = 4
	updateScore(currentPlayer, roll) // Player 2 rolls
	if currentPlayer.score != 49 {
		t.Errorf("Expected Player 2 score to be 49, got %d", currentPlayer.score)
	}

	// Test that game handles pig rolls correctly
	pigRoll := 1
	if checkPig(pigRoll) {
		t.Errorf("Expected checkPig(%d) to be false, got true", pigRoll)
	}
}

// Test for Story 1: Player 1 holds at 10, Player 2 holds at 15
func TestStory1(t *testing.T) {
	p1HoldStrategy := 10
	p2HoldStrategy := 15
	p1Wins, p2Wins := simulateGamesWithStrategies(p1HoldStrategy, p2HoldStrategy, 10)

	if p1Wins + p2Wins != 10 {
		t.Errorf("Total games played should be 10, but got %d wins for Player 1 and %d wins for Player 2", p1Wins, p2Wins)
	}
}

// Test for Story 2: Player 1 holds at 21, Player 2 holds at 1 to 100 (excluding 21)
func TestStory2(t *testing.T) {
	p1HoldStrategy := 21
	for p2HoldStrategy := 1; p2HoldStrategy <= 100; p2HoldStrategy++ {
		if p2HoldStrategy == 21 {
			continue // Skip the same strategy
		}
		p1Wins, p2Wins := simulateGamesWithStrategies(p1HoldStrategy, p2HoldStrategy, 10)

		if p1Wins + p2Wins != 10 {
			t.Errorf("Total games played should be 10 for Player 1 holding at 21 and Player 2 holding at %d, but got %d wins for Player 1 and %d wins for Player 2", p2HoldStrategy, p1Wins, p2Wins)
		}
	}
}

// Test for Story 3: Player 1 holds at 1 to 100, Player 2 holds at 1 to 100 (excluding same strategies)
func TestStory3(t *testing.T) {
	for p1HoldStrategy := 1; p1HoldStrategy <= 100; p1HoldStrategy++ {
		for p2HoldStrategy := 1; p2HoldStrategy <= 100; p2HoldStrategy++ {
			if p1HoldStrategy == p2HoldStrategy {
				continue // Skip the same strategy
			}
			p1Wins, p2Wins := simulateGamesWithStrategies(p1HoldStrategy, p2HoldStrategy, 10)

			if p1Wins + p2Wins != 10 {
				t.Errorf("Total games played should be 10 for Player 1 holding at %d and Player 2 holding at %d, but got %d wins for Player 1 and %d wins for Player 2", p1HoldStrategy, p2HoldStrategy, p1Wins, p2Wins)
			}
		}
	}
}