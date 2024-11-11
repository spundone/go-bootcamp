package main

import "fmt"

// NumberPredicate is a function type that takes an int and returns bool
type NumberPredicate func(int) bool

// FilterNumbers applies any of the conditions to the input slice
func FilterNumbers(numbers []int, conditions ...NumberPredicate) []int {
    var result []int
    
    numberLoop:
    for _, num := range numbers {
        // Check if any condition is true for the number
        for _, condition := range conditions {
            if condition(num) {
                result = append(result, num)
                continue numberLoop // Move to next number once we find a match
            }
        }
    }
    return result
}

// Predefined conditions
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func isMultipleOf(x int) NumberPredicate {
    return func(n int) bool {
        return n%x == 0
    }
}

func greaterThan(x int) NumberPredicate {
    return func(n int) bool {
        return n > x
    }
}

func lessThan(x int) NumberPredicate {
    return func(n int) bool {
        return n < x
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

    // Test case 1: prime, greater than 15, multiple of 5
    result1 := FilterNumbers(numbers, 
        isPrime,
        greaterThan(15),
        isMultipleOf(5),
    )
    fmt.Println("Result 1:", result1) 
    // Output: [2 3 5 7 10 11 13 15 16 17 18 19 20]

    // Test case 2: less than 6, multiple of 3
    result2 := FilterNumbers(numbers,
        lessThan(6),
        isMultipleOf(3),
    )
    fmt.Println("Result 2:", result2) 
    // Output: [1 2 3 4 5 6 9 12 15 18]
}
