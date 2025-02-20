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

	// Parse hold strategy ranges or single values
	p1HoldRange := strings.Split(p1HoldStrategyStr, "-")
	p2HoldRange := strings.Split(p2HoldStrategyStr, "-")

	var p1Start, p1End, p2Start, p2End int
	var err1, err2 error

	if len(p1HoldRange) == 1 {
		p1Start, err1 = strconv.Atoi(strings.TrimSpace(p1HoldRange[0]))
		p1End = p1Start // Single value, so start and end are the same
	} else {
		p1Start, p1End, err1 = parseHoldRange(p1HoldRange)
	}

	if len(p2HoldRange) == 1 {
		p2Start, err2 = strconv.Atoi(strings.TrimSpace(p2HoldRange[0]))
		p2End = p2Start // Single value, so start and end are the same
	} else {
		p2Start, p2End, err2 = parseHoldRange(p2HoldRange)
	}

	if err1 != nil || err2 != nil {
		fmt.Println("Hold strategies must be in the format 'start-end' with positive integers.")
		return
	}

	// Story 1: Single values for both players
	if p1End == p1Start && p2End == p2Start {
		p1Wins, p2Wins := simulateGamesWithStrategies(p1Start, p2Start, 10)
		fmt.Printf("Holding at %d vs Holding at %d: wins: %d/10 (%.1f%%), losses: %d/10 (%.1f%%)\n",
			p1Start, p2Start, p1Wins, float64(p1Wins)/10*100, p2Wins, float64(p2Wins)/10*100)
		return
	}

	// Story 2: Fixed value for p1 (e.g., 21) and range for p2 (1-100)
	if p1End == p1Start && p2End > p2Start {
		for k := 1; k <= 100; k++ {
			if k == p1Start {
				continue // Skip when strategies are the same
			}
			p1Wins, p2Wins := simulateGamesWithStrategies(p1Start, k, 10)
			fmt.Printf("Holding at %d vs Holding at %d: wins: %d/10 (%.1f%%), losses: %d/10 (%.1f%%)\n",
				p1Start, k, p1Wins, float64(p1Wins)/10*100, p2Wins, float64(p2Wins)/10*100)
		}
		return
	}

	// Story 3: Range for both players (1-100)
	if p1End > p1Start && p2End > p2Start {
		for k := 1; k <= 100; k++ {
			totalWins := 0
			totalLosses := 0
			for j := 1; j <= 100; j++ {
				if j == k {
					continue // Skip when strategies are the same
				}
				p1Wins, p2Wins := simulateGamesWithStrategies(k, j, 10)
				totalWins += p1Wins
				totalLosses += p2Wins
			}
			fmt.Printf("Result: Wins, losses staying at k = %d: %d/990 (%.1f%%), %d/990 (%.1f%%)\n",
				k, totalWins, float64(totalWins)/990*100, totalLosses, float64(totalLosses)/990*100)
		}
		return
	}
}

// New helper function to parse hold strategy ranges
func parseHoldRange(rangeStr []string) (int, int, error) {
	if len(rangeStr) != 2 {
		return 0, 0, fmt.Errorf("invalid range")
	}
	start, err1 := strconv.Atoi(strings.TrimSpace(rangeStr[0]))
	end, err2 := strconv.Atoi(strings.TrimSpace(rangeStr[1]))
	if err1 != nil || err2 != nil || start < 1 || end > 100 || start > end {
		return 0, 0, fmt.Errorf("invalid range")
	}
	return start, end, nil
}
