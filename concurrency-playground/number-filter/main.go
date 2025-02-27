package main

import (
	"fmt"
	"sync"
)

// evenNumbers filters out even numbers from a slice of integers.
func evenNumbers(num []int) []int {
	var evens []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, n := range num {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if n%2 == 0 {
				mu.Lock()
				evens = append(evens, n)
				mu.Unlock()
			}
		}(n)
	}
	wg.Wait()
	return evens
}

// oddNumbers filters out odd numbers from a slice of integers.
func oddNumbers(num []int) []int {
	var odds []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, n := range num {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if n%2 != 0 {
				mu.Lock()
				odds = append(odds, n)
				mu.Unlock()
			}
		}(n)
	}
	wg.Wait()
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
	var primes []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, n := range num {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if isPrime(n) {
				mu.Lock()
				primes = append(primes, n)
				mu.Unlock()
			}
		}(n)
	}
	wg.Wait()
	return primes
}

// oddPrimeNumbers filters out odd prime numbers from a slice of integers.
func oddPrimeNumbers(num []int) []int {
	var odds []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, n := range num {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if isPrime(n) && n%2 != 0 {
				mu.Lock()
				odds = append(odds, n)
				mu.Unlock()
			}
		}(n)
	}
	wg.Wait()
	return odds
}

// evenMultiplesOf5 filters out even multiples of 5 from a slice of integers.
func evenMultiplesOf5(num []int) []int {
	var evens []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, n := range num {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if n%2 == 0 && n%5 == 0 {
				mu.Lock()
				evens = append(evens, n)
				mu.Unlock()
			}
		}(n)
	}
	wg.Wait()
	return evens
}

// oddMultiplesOf3GreaterThan10 filters out odd multiples of 3 from a slice of integers greater than 10.
func oddMultiplesOf3GreaterThan10(num []int) []int {
	var odds []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, n := range num {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if n%2 != 0 && n%3 == 0 && n > 10 {
				mu.Lock()
				odds = append(odds, n)
				mu.Unlock()
			}
		}(n)
	}
	wg.Wait()
	return odds
}

type Condition func(n int) bool

func odd(n int) bool               { return n%2 != 0 }
func even(n int) bool              { return !odd(n) }
func prime(n int) bool             { return isPrime(n) }
func greaterThanN(n int) Condition { return func(m int) bool { return m > n } }
func multiplesOf(n int) Condition  { return func(m int) bool { return m%n == 0 } }
func lessThanN(n int) Condition    { return func(m int) bool { return m < n } }

// andNumbers filters out numbers from a slice of integers that match all the conditions.
func andNumbers(num []int, conditions ...Condition) []int {
	var filtered []int

	for _, num := range num {
		matchesAll := true
		for _, condition := range conditions {
			if !condition(num) {
				matchesAll = false
			}
		}
		if matchesAll {
			filtered = append(filtered, num)
		}
	}
	return filtered
}

// orNumbers filters out numbers from a slice of integers that match any of the conditions.
func orNumbers(num []int, conditions ...Condition) []int {
	var filtered []int
	for _, num := range num {
		for _, condition := range conditions {
			if condition(num) {
				filtered = append(filtered, num)
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
	fmt.Println("Even numbers:", evenNumbers(numbers))
	fmt.Println("Odd numbers:", oddNumbers(numbers))
	fmt.Println("Prime numbers:", primeNumbers(numbers))
	fmt.Println("Odd prime numbers:", oddPrimeNumbers(numbers))
	fmt.Println("Even multiples of 5:", evenMultiplesOf5(numbersExtended))
	fmt.Println("Odd multiples of 3 greater than 10", oddMultiplesOf3GreaterThan10(numbersExtended))
	fmt.Println("AndNumber", andNumbers(numbersExtended, odd, greaterThanN(5), multiplesOf(3)))
	fmt.Println("AndNumber2", andNumbers(numbersExtended, even, lessThanN(15), multiplesOf(3)))
	fmt.Println("AndNumber3", andNumbers(numbersExtended, prime, greaterThanN(10), lessThanN(15)))
	fmt.Println("OrNumber", orNumbers(numbersExtended, greaterThanN(15), multiplesOf(5)))
	fmt.Println("OrNumber2", orNumbers(numbersExtended, lessThanN(6), multiplesOf(3)))
	fmt.Println("OrNumber3", orNumbers(numbersExtended, greaterThanN(10), lessThanN(15)))

}
