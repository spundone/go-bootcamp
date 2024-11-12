# Number Filtering in Go

This project implements various number filtering exercises in Go, demonstrating different approaches to filtering numbers based on specific criteria.

## Project Structure

The project contains multiple implementations, each building on the previous one:

1. Even Numbers (`1/even.go`)
2. Odd Numbers (`2/odd.go`)
3. Prime Numbers (`3/prime.go`)
4. Odd Prime Numbers (`4/oddprime.go`)
5. Even Multiples of 5 (`5/evenmultiples5.go`)
6. Complex Conditions (`6/s6.go`)
7. Multiple Conditions with AND (`7/s7.go`)
8. Multiple Conditions with OR (`8/s8.go`)

## Implementations

### Basic Filters
- Even numbers filter
```go
startLine: 6
endLine: 14
```

- Odd numbers filter
```go
startLine: 6
endLine: 14
```

- Prime numbers filter
```go
startLine: 6
endLine: 27
```

### Advanced Filters

#### Function Types and Predicates
The advanced implementations use function types for flexible filtering:

```go
type NumberPredicate func(int) bool
```

#### Predefined Conditions
- isPrime
- isMultipleOf
- greaterThan
- lessThan
- isOdd
- isEven

## Usage Examples

1. Basic filtering:
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evenNumbers := filterEvenNumbers(numbers)
// Output: [2 4 6 8 10]
```

2. Advanced filtering with multiple conditions:
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
result := FilterNumbers(numbers,
    isPrime,
    greaterThan(5),
    isMultipleOf(3),
)
```
## Prerequisites

- Go 1.21 or higher
- Basic understanding of Go programming concepts

## Running the Examples

Each implementation can be run independently:

```bash
# Run even numbers filter
cd 1
go run even.go
```

```bash
# Run prime numbers filter
cd ../3
go run prime.go
```

```bash
# Run advanced filtering
cd ../8
go run s8.go
```

## Implementation Details

The project demonstrates several Go programming concepts:

1. Basic loops and conditionals
2. Function types and closures
3. Slices and slice operations
4. Higher-order functions
5. Modular code organization

## Contributing

Feel free to contribute by:
1. Forking the repository
2. Creating your feature branch
3. Committing your changes
4. Opening a pull request

## License

MIT License

## Acknowledgments

- Inspired by [one2n.io Go Bootcamp](https://one2n.io/go-bootcamp)
- Part of the Basic Number Filtering exercise
