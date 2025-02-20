package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Player struct {
	holdStrategy int // Hold at this value maybe 10 or 15
	score        int // Total score
}

// maybe redundant rolls a 6 sided die
func rollDie() int {
	return rand.Intn(6) + 1
}

// test die that always rolls 1
func checkPig(roll int) bool {
	return roll == 1
}

// Update the player's score by adding the rolled points
func updateScore(p *Player, points int) int {
	p.score += points
	return p.score
}

// Function to handle a single roll and its result
func handleRoll(roll int, turnTotal *int) bool {
	if roll == 1 {
		return false // Indicates that the player loses all points for this turn
	}
	*turnTotal += roll
	return true // Indicates a successful roll
}

func playTurn(p *Player) int {
	turnTotal := 0

	for {
		roll := rollDie()
		if !handleRoll(roll, &turnTotal) {
			return 0 // Player loses all points for this turn
		}

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

// New function to play the game n times and return the results
func simulateGames(p1, p2 *Player, n int) (int, int) {
	p1Wins := 0
	p2Wins := 0

	for i := 0; i < n; i++ {
		winner := playGame(p1, p2)
		if winner == 1 {
			p1Wins++
		} else {
			p2Wins++
		}
	}
	return p1Wins, p2Wins
}

func main() {
	// Parse the arguments here and pass them to simulateGames()
	p1HoldStrategy, _ := strconv.Atoi(os.Args[1]) // Convert to int
	p2HoldStrategy, _ := strconv.Atoi(os.Args[2]) // Convert to int

	p1 := &Player{holdStrategy: p1HoldStrategy}
	p2 := &Player{holdStrategy: p2HoldStrategy}

	p1Wins, p2Wins := simulateGames(p1, p2, 1000) // Simulate 1000 games
	fmt.Printf("Player 1 wins: %d, Player 2 wins: %d\n", p1Wins, p2Wins)
}
