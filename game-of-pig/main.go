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

type GameResult struct {
	p1Wins int
	p2Wins int
	games  int
}

// Rolls a 6 sided die
func rollDie() int {
	return rand.Intn(6) + 1
}

// Handles condition when dice rolls 1
func handlePigRoll(roll int, turnTotal *int) bool {
	if roll == 1 {
		return false // Indicates that the player loses all points for this turn
	}
	*turnTotal += roll
	return true // Indicates a successful roll
}

// Handles player switching on reaching hold value
func playTurn(p *Player) int {
	turnTotal := 0
	for {
		roll := rollDie()
		if !handlePigRoll(roll, &turnTotal) {
			return 0 // Player loses all points for this turn
		}
		if turnTotal >= p.holdStrategy {
			return turnTotal
		}
	}
}

// Helper function to validate hold strategies from input
func validateHoldStrategies(p1HoldNum, p2HoldNum int) error {
	if p1HoldNum <= 0 || p2HoldNum <= 0 {
		return fmt.Errorf("hold strategies must be greater than 0")
	}
	if p1HoldNum > 100 || p2HoldNum > 100 {
		return fmt.Errorf("hold strategies must not exceed 100")
	}
	return nil
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

// Simulates n no of games with holding strategy mentioned
func simulateGames(p1HoldNum, p2HoldNum int, numGames int) (GameResult, error) {
	if err := validateHoldStrategies(p1HoldNum, p2HoldNum); err != nil {
		return GameResult{}, err
	}

	p1 := &Player{holdStrategy: p1HoldNum}
	p2 := &Player{holdStrategy: p2HoldNum}
	result := GameResult{games: numGames}

	for i := 0; i < numGames; i++ {
		if playGame(p1, p2) == 1 {
			result.p1Wins++
		} else {
			result.p2Wins++
		}
	}
	return result, nil
}

// Simulates a game with fixed hold strategies for both players | Story1
func fixedHold(hold1, hold2 int) {
	result, err := simulateGames(hold1, hold2, 10)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
		hold1, hold2, result.p1Wins, result.games,
		float64(result.p1Wins)/float64(result.games)*100,
		result.p2Wins, result.games,
		float64(result.p2Wins)/float64(result.games)*100)
}

// Simulates games with one fixed hold and one range hold | Story2
func rangeHold1(fixedHold, rangeStart, rangeEnd int, rangeIsFirst bool) {
	for hold := rangeStart; hold <= rangeEnd; hold++ {
		if hold == fixedHold {
			continue // Skip when hold strategies are the same
		}
		var result GameResult
		if rangeIsFirst { // Bool to switch the player 1 and 2 and reuse existing code for story 2
			result, _ = simulateGames(hold, fixedHold, 10)
			fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
				hold, fixedHold,
				result.p1Wins, result.games,
				float64(result.p1Wins)/float64(result.games)*100,
				result.p2Wins, result.games,
				float64(result.p2Wins)/float64(result.games)*100)
		} else {
			result, _ = simulateGames(fixedHold, hold, 10)
			fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
				fixedHold, hold,
				result.p1Wins, result.games,
				float64(result.p1Wins)/float64(result.games)*100,
				result.p2Wins, result.games,
				float64(result.p2Wins)/float64(result.games)*100)
		}
	}
}

// Simulates games with both players using range holds | Story3
func rangeHold2(range1Start, range1End, range2Start, range2End int) {
	for hold1 := range1Start; hold1 <= range1End; hold1++ {
		totalWins := 0
		totalGames := 0
		for hold2 := range2Start; hold2 <= range2End; hold2++ {
			if hold2 == hold1 {
				continue // Skip when hold strategies are the same
			}
			result, _ := simulateGames(hold1, hold2, 10)
			totalWins += result.p1Wins
			totalGames += result.games
		}
		fmt.Printf("Result: Wins, losses staying at hold = %d: %d/%d (%.1f%%), %d/%d (%.1f%%)\n",
			hold1, totalWins, totalGames,
			float64(totalWins)/float64(totalGames)*100,
			totalGames-totalWins, totalGames,
			float64(totalGames-totalWins)/float64(totalGames)*100)
	}
}

// Input parsing for range
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

// Temp parser
func parseInput(args []string) (p1Start, p1End, p2Start, p2End int, err error) {
	if len(args) < 2 {
		return 0, 0, 0, 0, fmt.Errorf("insufficient arguments")
	}

	p1Range := strings.Split(args[0], "-")
	p2Range := strings.Split(args[1], "-")

	if len(p1Range) == 1 {
		p1Start, err = strconv.Atoi(strings.TrimSpace(p1Range[0]))
		p1End = p1Start
	} else {
		p1Start, p1End, err = parseHoldRange(p1Range)
	}

	if err != nil {
		return 0, 0, 0, 0, err
	}

	if len(p2Range) == 1 {
		p2Start, err = strconv.Atoi(strings.TrimSpace(p2Range[0]))
		p2End = p2Start
	} else {
		p2Start, p2End, err = parseHoldRange(p2Range)
	}
	return p1Start, p1End, p2Start, p2End, err
}

func main() {
	p1Start, p1End, p2Start, p2End, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Println("Usage: ./pig <player1_hold_strategy> <player2_hold_strategy>")
		fmt.Println("Error:", err)
		return
	}

	// Determine which func to run based on input
	if p1End == p1Start && p2End == p2Start {
		fixedHold(p1Start, p2Start)
	} else if p1End > p1Start && p2End == p2Start {
		rangeHold1(p2Start, p1Start, p1End, true) // First player uses range like 1-100 10
	} else if p1End == p1Start && p2End > p2Start {
		rangeHold1(p1Start, p2Start, p2End, false) // Second player uses range like 21 1-100
	} else if p1End > p1Start && p2End > p2Start {
		rangeHold2(p1Start, p1End, p2Start, p2End)
	}
}
