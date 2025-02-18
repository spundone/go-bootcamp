package main

import (
	"fmt"
)


type Condition func(n int) bool

func odd(n int) bool               { return n%2 != 0 }
func even(n int) bool              { return !odd(n) }
func prime(n int) bool             { return isPrime(n) }
func greaterThanN(n int) Condition { return func(m int) bool { return m > n } }
func multiplesOf(n int) Condition  { return func(m int) bool { return m%n == 0 } }
func lessThanN(n int) Condition    { return func(m int) bool { return m < n } }

// filterEven filters out even numbers from a slice of integers.
func filterEven(num []int) []int {
	if len(num) == 0 {
		return []int{}
	}
	var evens []int
	for _, n := range num {
		if even(n) {
			evens = append(evens, n)
		}
	}
	return evens
}

// filterOdd filters out odd numbers from a slice of integers.
func filterOdd(num []int) []int {
	if len(num) == 0 {
		return []int{}
	}
	var odds []int
	for _, n := range num {
		if odd(n) {
			odds = append(odds, n)
		}
	}
	return odds
}

// isPrime checks if a number is prime.
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ { //another way is to only check up to sqrt(num)
		if num%i == 0 {
			return false
		}
	}
	return true
}

// primeNumbers filters out prime numbers from a slice of integers.
func primeNumbers(num []int) []int {
	if len(num) == 0 {
		return []int{}
	}
	var primes []int
	for _, n := range num {
		if isPrime(n) {
			primes = append(primes, n)
		}
	}
	return primes
}

// oddPrimeNumbers filters out odd prime numbers from a slice of integers.
func oddPrimeNumbers(num []int) []int {
	if len(num) == 0 {
		return []int{}
	}
	var odds []int
	for _, n := range num {
		if isPrime(n) && n%2 != 0 {
			odds = append(odds, n)
		}
	}
	return odds
}

// evenMultiplesOf5 filters out even multiples of 5 from a slice of integers.
func evenMultiplesOf5(num []int) []int {
	if len(num) == 0 {
		return []int{}
	}
	var evens []int
	for _, n := range num {
		if n%2 == 0 && n%5 == 0 {
			evens = append(evens, n)
		}
	}
	return evens
}

// oddMultiplesOf3GreaterThan10 filters out odd multiples of 3 from a slice of integers greater than 10.
func oddMultiplesOf3GreaterThan10(num []int) []int {
	if len(num) == 0 {
		return []int{}
	}
	var odds []int
	for _, n := range num {
		if n%2 != 0 && n%3 == 0 && n > 10 {
			odds = append(odds, n)
		}
	}
	return odds
}

// filtersAll filters out numbers from a slice of integers that match all the conditions.
func filtersAll(num []int, conditions ...Condition) []int {
	if len(num) == 0 {
		return []int{}
	}
	var filtered []int

	for _, n := range num {
		matchesAll := true
		for _, condition := range conditions {
			if !condition(n) {
				matchesAll = false
			}
		}
		if matchesAll {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

// filterAny filters out numbers from a slice of integers that match any of the conditions.
func filterAny(num []int, conditions ...Condition) []int {
	if len(num) == 0 {
		return []int{}
	}
	var filtered []int
	for _, n := range num {
		for _, condition := range conditions {
			if condition(n) {
				filtered = append(filtered, n)
				break
			}
		}
	}
	return filtered
}

// main function
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbersExtended := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("Even numbers:", filterEven(numbers))
	fmt.Println("Odd numbers:", filterOdd(numbers))
	fmt.Println("Prime numbers:", primeNumbers(numbers))
	fmt.Println("Odd prime numbers:", oddPrimeNumbers(numbers))
	fmt.Println("Even multiples of 5:", evenMultiplesOf5(numbersExtended))
	fmt.Println("Odd multiples of 3 greater than 10", oddMultiplesOf3GreaterThan10(numbersExtended))
	fmt.Println("AndNumber", filtersAll(numbersExtended, odd, greaterThanN(5), multiplesOf(3)))
	fmt.Println("AndNumber2", filtersAll(numbersExtended, even, lessThanN(15), multiplesOf(3)))
	fmt.Println("AndNumber3", filtersAll(numbersExtended, prime, greaterThanN(10), lessThanN(15)))
	fmt.Println("OrNumber", filterAny(numbersExtended, greaterThanN(15), multiplesOf(5)))
	fmt.Println("OrNumber2", filterAny(numbersExtended, lessThanN(6), multiplesOf(3)))
	fmt.Println("OrNumber3", filterAny(numbersExtended, greaterThanN(10), lessThanN(15)))

}
