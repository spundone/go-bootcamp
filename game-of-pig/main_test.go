package main

import (
	"fmt"
	"os"
	"testing"
)

func TestStory1(t *testing.T) {
	// Redirect stdout to capture the output
	stdout := os.Stdout
	defer func() { os.Stdout = stdout }()

	// Set the strategies for Player 1 and Player 2
	p1Strategy := 10
	p2Strategy := 15

	// Simulate the game
	p1 := &Player{holdStrategy: p1Strategy}
	p2 := &Player{holdStrategy: p2Strategy}

	p1Wins := 0
	p2Wins := 0

	// Play games until one player wins
	for p1Wins+p2Wins < 10 {
		winner := playGame(p1, p2)
		if winner == 1 {
			p1Wins++
		} else {
			p2Wins++
		}
	}

	// Check the output
	expectedOutput := fmt.Sprintf("Holding at %d vs Holding at %d: wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
		p1Strategy, p2Strategy, p1Wins, p1Wins+p2Wins, float64(p1Wins)*100/float64(p1Wins+p2Wins),
		p2Wins, p1Wins+p2Wins, float64(p2Wins)*100/float64(p1Wins+p2Wins))

	if buf.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, buf.String())
	}
}

func TestStory2(t *testing.T) {
	// Player 1 has a fixed strategy of holding at 21
	p1Strategy := 21

	// Simulate games for Player 2's strategies from 1 to 100, excluding 21
	for p2Strategy := 1; p2Strategy <= 100; p2Strategy++ {
		if p2Strategy == 21 {
			continue // Skip the strategy of holding at 21
		}

		p1 := &Player{holdStrategy: p1Strategy}
		p2 := &Player{holdStrategy: p2Strategy}

		p1Wins := 0
		p2Wins := 0

		// Simulate 10 games for each strategy of Player 2
		for i := 0; i < 10; i++ {
			winner := playGame(p1, p2)
			if winner == 1 {
				p1Wins++
			} else {
				p2Wins++
			}
		}

		// Print the results for the current strategy
		fmt.Printf("Holding at %d vs Holding at %d: wins: %d/10 (%.1f%%), losses: %d/10 (%.1f%%)\n",
			p1Strategy, p2Strategy, p1Wins, float64(p1Wins)*100/10,
			p2Wins, float64(p2Wins)*100/10)
	}
}

func TestStory3(t *testing.T) {
	// Iterate over Player 1's strategies from 1 to 100
	for p1Strategy := 1; p1Strategy <= 100; p1Strategy++ {
		p1Wins := 0
		p2Wins := 0

		// Iterate over Player 2's strategies from 1 to 100, excluding Player 1's strategy
		for p2Strategy := 1; p2Strategy <= 100; p2Strategy++ {
			if p2Strategy == p1Strategy {
				continue // Skip the strategy that matches Player 1's
			}

			p1 := &Player{holdStrategy: p1Strategy}
			p2 := &Player{holdStrategy: p2Strategy}

			// Simulate 10 games for each strategy combination
			for i := 0; i < 10; i++ {
				winner := playGame(p1, p2)
				if winner == 1 {
					p1Wins++
				} else {
					p2Wins++
				}
			}
		}

		// Total games played against Player 2's strategies
		totalGames := 990 // 99 strategies * 10 games each
		// Print the results for Player 1's current strategy
		fmt.Printf("Result: Wins, losses staying at k = %d: %d/%d (%.1f%%), %d/%d (%.1f%%)\n",
			p1Strategy, p1Wins, totalGames, float64(p1Wins)*100/float64(totalGames),
			totalGames-p1Wins, totalGames, float64(totalGames-p1Wins)*100/float64(totalGames))
	}
}
