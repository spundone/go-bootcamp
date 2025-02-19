# Game of Pig Implementation in Go

A Go implementation of the Game of Pig, a simple dice game where players take turns rolling a die to accumulate points while managing risk.

## Game Rules

1. Each turn, a player repeatedly rolls a die until either a 1 is rolled or the player decides to “hold”:

2. If the player rolls a 1, they score nothing, and it becomes the next player’s turn.

3. If the player rolls any other number, it is added to their turn total, and their turn continues. The player can then decide to continue to roll again or “hold”.

4. If a player chooses to “hold”, their current turn total is added to their score, and it becomes the next player’s turn.

5. First player to reach 100 points wins.

## Features

- Simulates dice rolling and turn management
- Implements different player strategies (hold at different values)
- Supports multiple game simulations
- Provides win rate analysis for different strategies
