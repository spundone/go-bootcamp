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
		{1, true},
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


