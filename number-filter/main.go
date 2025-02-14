package main

import (
    "fmt"
)

// evenNumbers filters out even numbers from a slice of integers.
func evenNumbers(num []int) []int {
    var evens []int
    for _, num := range num {
        if num%2 == 0 {
            evens = append(evens, num)
        }
    }
    return evens
}

// oddNumbers filters out odd numbers from a slice of integers.
func oddNumbers(num []int) []int {
    var odds []int
    for _, num := range num {
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


// primeNumbers filters out prime numbers from a slice of integers.
func primeNumbers(num []int) []int {
	var primes []int
	for _, num := range num {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}
	return primes
}

// oddPrimeNumbers filters out odd prime numbers from a slice of integers.
func oddPrimeNumbers(num []int) []int {
	var odds []int
	for _, num := range num {
		if isPrime(num) && num%2 != 0 {
			odds = append(odds, num)
		}
	}
	return odds
}

// evenMultiplesOf5 filters out even multiples of 5 from a slice of integers.
func evenMultiplesOf5(num []int) []int {
	var evens []int
	for _, num := range num {
		if num%2 == 0 && num%5 == 0 {
			evens = append(evens, num)
		}
	}
	return evens
}

// oddMultiplesOf3GreaterThan10 filters out odd multiples of 3 from a slice of integers greater than 10.
func oddMultiplesOf3GreaterThan10(num []int) []int {
	var odds []int
	for _, num := range num {
		if num%2 != 0 && num%3 == 0 && num > 10 {	
			odds = append(odds, num)
		}
	}
	return odds
}

// numbers filters out numbers from a slice of integers that match all the conditions.
func numbers(num []int, conditions []func(int) bool) []int {
	var filtered []int
	for _, num := range num {
		matchesAll := true
		for _, condition := range conditions {
			if !condition(num) {
				matchesAll = false
				break
			}
		}
		if matchesAll {
			filtered = append(filtered, num)
		}
	}
	return filtered
}

// main function
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbersExtended := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("Even numbers:", evenNumbers(numbers))
	fmt.Println("Odd numbers:", oddNumbers(numbers))
	fmt.Println("Prime numbers:", primeNumbers(numbers))
	fmt.Println("Odd prime numbers:", oddPrimeNumbers(numbers))
	fmt.Println("Even multiples of 5:", evenMultiplesOf5(numbersExtended))
	fmt.Println("Odd multiples of 3 greater than 10", oddMultiplesOf3GreaterThan10(numbersExtended))
	//fmt.Println("Even multiples of 5 and odd multiples of 3", numbers(numbersExtended, []func(int) bool{evenMultiplesOf5, oddMultiplesOf3}))
}
