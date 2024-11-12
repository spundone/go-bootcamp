package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Player struct {
	strategy int // Hold at this value
	score    int // Total score
}

// rollDie simulates rolling a 6-sided die
func rollDie() int {
	return rand.Intn(6) + 1
}

// playTurn simulates a single turn for a player
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

// simulateGames plays multiple games and returns wins for player 1
func simulateGames(p1Strategy, p2Strategy, numGames int) int {
	p1 := &Player{strategy: p1Strategy}
	p2 := &Player{strategy: p2Strategy}
	p1Wins := 0

	for i := 0; i < numGames; i++ {
		winner := playGame(p1, p2)
		if winner == 1 {
			p1Wins++
		}
	}

	return p1Wins
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) != 3 {
		fmt.Println("Usage: ./pig <strategy1> <strategy2>")
		os.Exit(1)
	}

	// Parse first argument
	if os.Args[1] == "1-100" {
		// Story 3: Run all strategies for player 1
		runStory3()
		return
	}

	p1Strategy, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid strategy for player 1: %s\n", os.Args[1])
		os.Exit(1)
	}

	// Parse second argument
	if os.Args[2] == "1-100" {
		// Story 2: Run all strategies for player 2
		runStory2(p1Strategy)
		return
	}

	p2Strategy, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid strategy for player 2: %s\n", os.Args[2])
		os.Exit(1)
	}

	// Story 1: Single match between two strategies
	wins := simulateGames(p1Strategy, p2Strategy, 10)
	losses := 10 - wins
	fmt.Printf("Holding at %d vs Holding at %d: wins: %d/10 (%.1f%%), losses: %d/10 (%.1f%%)\n",
		p1Strategy, p2Strategy, wins, float64(wins)*10, losses, float64(losses)*10)
}

func runStory2(p1Strategy int) {
	for i := 1; i <= 100; i++ {
		if i == p1Strategy {
			continue // Skip when strategies are the same
		}
		wins := simulateGames(p1Strategy, i, 10)
		losses := 10 - wins
		fmt.Printf("Holding at %d vs Holding at %d: wins: %d/10 (%.1f%%), losses: %d/10 (%.1f%%)\n",
			p1Strategy, i, wins, float64(wins)*10, losses, float64(losses)*10)
	}
}

func runStory3() {
	for k := 1; k <= 100; k++ {
		totalGames := 0
		totalWins := 0

		for opponent := 1; opponent <= 100; opponent++ {
			if opponent == k {
				continue // Skip when strategies are the same
			}
			wins := simulateGames(k, opponent, 10)
			totalWins += wins
			totalGames += 10
		}

		losses := totalGames - totalWins
		winRate := float64(totalWins) * 100 / float64(totalGames)
		lossRate := float64(losses) * 100 / float64(totalGames)

		fmt.Printf("Result: Wins, losses staying at k = %d: %d/%d (%.1f%%), %d/%d (%.1f%%)\n",
			k, totalWins, totalGames, winRate, losses, totalGames, lossRate)
	}
}
