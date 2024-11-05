package main

import "fmt"

// Function to filter odd numbers from a list
func filterOddNumbers(numbers []int) []int {
    var oddNumbers []int
    for _, num := range numbers {
        if num % 2 != 0 {
            oddNumbers = append(oddNumbers, num)
        }
    }
    return oddNumbers
}

func main() {
    // Example list of integers
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Filter and print even numbers
    oddNumbers := filterOddNumbers(numbers)
    fmt.Println("Odd numbers:", oddNumbers)
}