package main

import "fmt"

// Function to filter prime numbers from a list
func filterPrimeNumbers(numbers []int) []int {
    var primeNumbers []int
    for _, num := range numbers {
        if num > 1 && isPrime(num) {
            primeNumbers = append(primeNumbers, num)
        }
    }
    return primeNumbers
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
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Filter and print prime numbers
    primeNumbers := filterPrimeNumbers(numbers)
    fmt.Println("Prime numbers:", primeNumbers)
}
