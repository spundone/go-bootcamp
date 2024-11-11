package main

import "fmt"

// Function to filter even numbers from a list
func filterEvenNumbers(numbers []int) []int {
    var evenNumbers []int
    for _, num := range numbers {
        if num % 2 == 0 {
            evenNumbers = append(evenNumbers, num)
        }
    }
    return evenNumbers
}

func main() {
    // Example list of integers
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Filter and print even numbers
    evenNumbers := filterEvenNumbers(numbers)
    fmt.Println("Even numbers:", evenNumbers)
}