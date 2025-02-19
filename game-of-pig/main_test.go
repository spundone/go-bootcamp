package main

import (
	"testing"
)

func TestRollDie(t *testing.T) {
	// Test that die rolls are within valid range
	for i := 0; i < 1000; i++ { // Test 1000 times
		roll := rollDie()
		if roll < 1 || roll > 6 {
			t.Errorf("Invalid die roll: %d", roll)
		}
	}
}

