package main

import "fmt"

// NumberPredicate is a function type that takes an int and returns bool
type NumberPredicate func(int) bool

// FilterNumbers applies all conditions to the input slice
func FilterNumbers(numbers []int, conditions ...NumberPredicate) []int {
    var result []int
    
    numberLoop:
    for _, num := range numbers {
        // Check all conditions for each number
        for _, condition := range conditions {
            if !condition(num) {
                continue numberLoop
            }
        }
        result = append(result, num)
    }
    return result
}

// Predefined conditions
func isOdd(n int) bool {
    return n%2 != 0
}

func isEven(n int) bool {
    return n%2 == 0
}

func isMultipleOf3(n int) bool {
    return n%3 == 0
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

    // Test case 1: odd, greater than 5, multiple of 3
    result1 := FilterNumbers(numbers, 
        isOdd,
        greaterThan(5),
        isMultipleOf3,
    )
    fmt.Println("Result 1:", result1) // Output: [9 15]

    // Test case 2: even, less than 15, multiple of 3
    result2 := FilterNumbers(numbers,
        isEven,
        lessThan(15),
        isMultipleOf3,
    )
    fmt.Println("Result 2:", result2) // Output: [6 12]
}
