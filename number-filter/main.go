package main

import (
    "fmt"
)

// filterEvenNumbers filters out even numbers from a slice of integers.
func filterEvenNumbers(numbers []int) []int {
    var evens []int
    for _, num := range numbers {
        if num%2 == 0 {
            evens = append(evens, num)
        }
    }
    return evens
}

// filterOddNumbers filters out odd numbers from a slice of integers.
func filterOddNumbers(numbers []int) []int {
    var odds []int
    for _, num := range numbers {
        if num%2 != 0 {
            odds = append(odds, num)
        }
    }
    return odds
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Even numbers:", filterEvenNumbers(numbers))
	fmt.Println("Odd numbers:", filterOddNumbers(numbers))
}
