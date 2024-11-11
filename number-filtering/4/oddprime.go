package main

import "fmt"

// Function to filter odd prime numbers from a list
func filterOddPrimeNumbers(numbers []int) []int {
    var oddPrimeNumbers []int
    for _, num := range numbers {
        if num > 1 && num%2 != 0 && isPrime(num) {
            oddPrimeNumbers = append(oddPrimeNumbers, num)
        }
    }
    return oddPrimeNumbers
}

// Function to check if a number is prime
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

func main() {
    // Example list of integers
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
    
    // Filter and print odd prime numbers
    oddPrimeNumbers := filterOddPrimeNumbers(numbers)
    fmt.Println("Odd prime numbers:", oddPrimeNumbers)
}
