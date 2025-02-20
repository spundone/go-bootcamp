package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

// New function to simulate multiple games with different strategies
func simulateGamesWithStrategies(p1HoldStrategy int, p2HoldStrategy int, n int) (int, int) {
	p1Wins := 0
	p2Wins := 0

	for i := 0; i < n; i++ {
		p1 := &Player{holdStrategy: p1HoldStrategy}
		p2 := &Player{holdStrategy: p2HoldStrategy}
		winner := playGame(p1, p2)
		if winner == 1 {
			p1Wins++
		} else {
			p2Wins++
		}
	}
	return p1Wins, p2Wins
}

// Main function to handle command-line arguments and simulate games
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./pig <player1_hold_strategy> <player2_hold_strategy>")
		return
	}

	p1HoldStrategyStr := os.Args[1]
	p2HoldStrategyStr := os.Args[2]

	// Check for empty values or spaces
	if p1HoldStrategyStr == "" || p2HoldStrategyStr == "" {
		fmt.Println("Hold strategies cannot be empty.")
		return
	}

	// Convert to int and check for valid values
	p1HoldStrategy, err1 := strconv.Atoi(strings.TrimSpace(p1HoldStrategyStr))
	p2HoldStrategy, err2 := strconv.Atoi(strings.TrimSpace(p2HoldStrategyStr))

	if err1 != nil || err2 != nil || p1HoldStrategy <= 0 || p2HoldStrategy <= 0 {
	 	fmt.Println("Hold strategies must be positive integers greater than 0.")
	 	return
	}

	// Simulate 10 games between the two players
	p1Wins, p2Wins := simulateGamesWithStrategies(p1HoldStrategy, p2HoldStrategy, 10) // Simulate 10 games
	fmt.Printf("Holding at %d vs Holding at %d: wins: %d/10 (%.1f%%), losses: %d/10 (%.1f%%)\n",
		p1HoldStrategy, p2HoldStrategy, p1Wins, float64(p1Wins)/10*100, p2Wins, float64(p2Wins)/10*100)
}
