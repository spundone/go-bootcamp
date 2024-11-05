package main

import "fmt"

func filterNumbers(numbers []int) []int {
    var result []int
    for _, num := range numbers {
        // Check if number is odd (num % 2 != 0)
        // AND multiple of 3 (num % 3 == 0)
        // AND greater than 10
        if num%2 != 0 && num%3 == 0 && num > 10 {
            result = append(result, num)
        }
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
    filtered := filterNumbers(numbers)
    fmt.Println(filtered) // Output: [15]
}
