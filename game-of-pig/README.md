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

## Implementation Details

The core game logic is implemented in `gop.go`:

- Player structure and game state management:
```go:game-of-pig/gop.go
startLine: 12
endLine: 15

- Turn simulation with strategy-based decision making:
```go:game-of-pig/gop.go
startLine: 22
endLine: 37
```

## Usage

1. Run a single game between two strategies:
```bash
./gop <strategy1> <strategy2>
```

2. Test strategy against all possible opponent strategies:
```bash
./gop <strategy> 1-100
```

3. Find optimal strategy by testing all combinations:
```bash
./gop 1-100 1-100
```

## Example Output

```
Holding at 25 vs Holding at 20: wins: 6/10 (60.0%), losses: 4/10 (40.0%)
```

## Testing

The project includes comprehensive tests in `gop_test.go`:

- Die rolling validation
- Turn mechanics testing
- Game simulation verification
- Strategy effectiveness testing

Run tests using:
```bash
go test
```

## Project Structure

```
game-of-pig/
├── gop.go         # Main implementation
├── gop_test.go    # Test suite
└── README.md      # Documentation
```

## Prerequisites

- Go 1.21 or higher
- Basic understanding of probability and game theory

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License.

## Acknowledgments

- Inspired by [one2nc/gameofpig](https://github.com/one2nc/gameofpig)
- Part of the [one2n.io Go Bootcamp](https://one2n.io/go-bootcamp/go-projects/a-game-of-pig) course