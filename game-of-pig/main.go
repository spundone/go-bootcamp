package main

import (
	// "fmt"
	"math/rand"
)

type Player struct {
	holdStrategy int // Hold at this value maybe 10 or 15
	score        int // Total score
}

// maybe redundant rolls a 6 sided die
func rollDie() int {
	return rand.Intn(6) + 1
}

func playTurn(p *Player) int {
	turnTotal := 0

	for {
		roll := rollDie()
		if roll == 1 {
			return 0
		}

		turnTotal += roll
		if turnTotal >= p.holdStrategy {
			return turnTotal
		}
	}
}

// playGame simulates a complete game between two players
func playGame(p1, p2 *Player) int {
	p1.score = 0
	p2.score = 0

	for {
		// Player 1's turn
		p1.score += playTurn(p1)
		if p1.score >= 100 {
			return 1 // Player 1 wins
		}

		// Player 2's turn
		p2.score += playTurn(p2)
		if p2.score >= 100 {
			return 2 // Player 2 wins
		}
	}
}

func simulateGames(p1Strategy, p2Strategy, numGames int) (int, int) {
	p1 := &Player{holdStrategy: p1Strategy}
	p2 := &Player{holdStrategy: p2Strategy}

	p1Wins := 0
	p2Wins := 0

	for i := 0; i < numGames; i++ {
		winner := playGame(p1, p2)
		if winner == 1 {
			p1Wins++
		} else {
			p2Wins++
		}
	}

	return p1Wins, p2Wins
}

