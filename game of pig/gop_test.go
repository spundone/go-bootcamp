package main

import (
	"testing"
)

func TestRollDie(t *testing.T) {
	// Test that die rolls are within valid range
	for i := 0; i < 1000; i++ {
		roll := rollDie()
		if roll < 1 || roll > 6 {
			t.Errorf("Invalid die roll: %d", roll)
		}
	}
}

func TestPlayTurn(t *testing.T) {
	tests := []struct {
		name     string
		strategy int
		maxScore int // Maximum possible score in a single turn
	}{
		{"Conservative strategy", 10, 30},
		{"Aggressive strategy", 25, 50},
		{"Very conservative strategy", 5, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{strategy: tt.strategy}
			for i := 0; i < 100; i++ {
				score := playTurn(p)
				if score < 0 || score > tt.maxScore {
					t.Errorf("Invalid turn score: %d", score)
				}
				if score > 0 && score < tt.strategy {
					t.Errorf("Score %d is less than strategy %d", score, tt.strategy)
				}
			}
		})
	}
}

func TestPlayGame(t *testing.T) {
	tests := []struct {
		name        string
		p1Strategy  int
		p2Strategy  int
		numGames    int
		minWinRate  float64 // Minimum expected win rate for either player
		maxWinRate  float64 // Maximum expected win rate for either player
	}{
		{
			name:       "Equal strategies",
			p1Strategy: 20,
			p2Strategy: 20,
			numGames:   100,
			minWinRate: 0.3, // Expect roughly equal wins
			maxWinRate: 0.7,
		},
		{
			name:       "Conservative vs Aggressive",
			p1Strategy: 10,
			p2Strategy: 25,
			numGames:   100,
			minWinRate: 0.2, // One strategy might be better
			maxWinRate: 0.8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1Wins := 0
			for i := 0; i < tt.numGames; i++ {
				p1 := &Player{strategy: tt.p1Strategy}
				p2 := &Player{strategy: tt.p2Strategy}
				winner := playGame(p1, p2)
				if winner == 1 {
					p1Wins++
				}
				// Verify final scores
				if p1.score < 100 && p2.score < 100 {
					t.Error("Game ended without a winner reaching 100")
				}
			}
			winRate := float64(p1Wins) / float64(tt.numGames)
			if winRate < tt.minWinRate || winRate > tt.maxWinRate {
				t.Errorf("Win rate %.2f outside expected range [%.2f, %.2f]",
					winRate, tt.minWinRate, tt.maxWinRate)
			}
		})
	}
} 