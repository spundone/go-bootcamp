package main

import (
	"testing"
)

func TestRollDie(t *testing.T) { //unnecessary imo
	// Test that die rolls are within valid range
	for i := 0; i < 1000; i++ { // Test 1000 times
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
			p := &Player{holdStrategy: tt.strategy}
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

