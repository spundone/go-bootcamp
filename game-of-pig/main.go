package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	strategy int // Hold at this value
	score    int // Total score
}

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
		if turnTotal >= p.strategy {
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


func main() {
	fmt.Println(rand.Intn(6) + 1)
	// fmt.Println(rollDie)
}