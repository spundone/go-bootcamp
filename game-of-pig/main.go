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

// Turn and game logic
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

// Simulates n no of games with holding strategy mentioned
func simulateGames(p1HoldNum, p2HoldNum int, numGames int) GameResult {
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
	return result
}

// Will run story 1 with 2 int parameters passed
func runStory1(p1HoldNum, p2HoldNum int) {
	result := simulateGames(p1HoldNum, p2HoldNum, 10)
	fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
		p1HoldNum, p2HoldNum, result.p1Wins, result.games,
		float64(result.p1Wins)/float64(result.games)*100,
		result.p2Wins, result.games,
		float64(result.p2Wins)/float64(result.games)*100)
}

// Will run story 2 with 1 int and 1 range parameter passed
func runStory2(p1HoldNum, p2HoldRangeStart, p2HoldRangeEnd int) {
	for k := p2HoldRangeStart; k <= p2HoldRangeEnd; k++ {
		if k == p1HoldNum {
			continue // Skips in case same holding strategy
		}
		result := simulateGames(p1HoldNum, k, 10)
		fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
			p1HoldNum, k, result.p1Wins, result.games,
			float64(result.p1Wins)/float64(result.games)*100,
			result.p2Wins, result.games,
			float64(result.p2Wins)/float64(result.games)*100)
	}
}

// Will run story 3 in case both parameters are range
func runStory3(p1HoldRangeStart, p1HoldRangeEnd, p2HoldRangeStart, p2HoldRangeEnd int) {
	for k := p1HoldRangeStart; k <= p1HoldRangeEnd; k++ {
		totalWins := 0
		totalLosses := 0
		for j := p2HoldRangeStart; j <= p2HoldRangeEnd; j++ {
			if j == k {
				continue // Skips in case holding same value
			}
			result := simulateGames(k, j, 10)
			totalWins += result.p1Wins
			totalLosses += result.p2Wins
		}
		fmt.Printf("Result: Wins, losses staying at k = %d: %d/990 (%.1f%%), %d/990 (%.1f%%)\n",
			k, totalWins, float64(totalWins)/990*100,// hardcoded to be fixed
			totalLosses, float64(totalLosses)/990*100)
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

	// Determine which story to run based on input
	if p1End == p1Start && p2End == p2Start {
		runStory1(p1Start, p2Start)
	} else if p1End == p1Start && p2End > p2Start {
		runStory2(p1Start, p2Start, p2End)
	} else if p1End > p1Start && p2End > p2Start {
		runStory3(p1Start, p1End, p2Start, p2End)
	}
}
