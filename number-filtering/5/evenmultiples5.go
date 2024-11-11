package main

import "fmt"

// Function to filter even multiples of 5 numbers from a list
func filterEvenMultiplesOf5(numbers []int) []int {
    var evenMultiplesOf5 []int
    for _, num := range numbers {
        if num % 5 == 0 && num % 2 == 0 {
            evenMultiplesOf5 = append(evenMultiplesOf5, num)
        }
    }
    return evenMultiplesOf5
}

func main() {
    // Example list of integers
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 25}
    
    // Filter and print prime numbers
    evenMultiplesOf5 := filterEvenMultiplesOf5(numbers)
    fmt.Println("Even multiples of 5:", evenMultiplesOf5)
}