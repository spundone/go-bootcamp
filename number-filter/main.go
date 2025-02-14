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

// isPrime checks if a number is prime.
func isPrime(num int) bool {
    if num <= 1 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}


// filterPrimeNumbers filters out prime numbers from a slice of integers.
func filterPrimeNumbers(numbers []int) []int {
	var primes []int
	for _, num := range numbers {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}
	return primes
}

// filterOddPrimeNumbers filters out odd prime numbers from a slice of integers.
func filterOddPrimeNumbers(numbers []int) []int {
	var odds []int
	for _, num := range numbers {
		if isPrime(num) && num%2 != 0 {
			odds = append(odds, num)
		}
	}
	return odds
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Even numbers:", filterEvenNumbers(numbers))
	fmt.Println("Odd numbers:", filterOddNumbers(numbers))
	fmt.Println("Prime numbers:", filterPrimeNumbers(numbers))
	fmt.Println("Odd prime numbers:", filterOddPrimeNumbers(numbers))
}
