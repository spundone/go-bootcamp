# Game of Pig Implementation in Go

A Go implementation of the Game of Pig, a simple dice game where players take turns rolling a die to accumulate points while managing risk.

## Game Rules

1. Each turn, a player repeatedly rolls a die until either:
   - Player rolls a 1 (loses turn, no points accumulated)
   - Player holds (keeps accumulated points)

2. First player to reach 100 points wins.

## Features

- Simulates dice rolling and turn management
- Implements different player strategies (hold at different values)
- Supports multiple game simulations
- Provides win rate analysis for different strategies
