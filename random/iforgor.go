// code 
package mathworks

import *math*

func evenNumbers (nums [] int) [] int {
	var output []int
	for _, n := range nums {
		if n%2 ==0 {
			output =append(output, n)
		}
	}
	return output
}

func oddNumbers (nums [] int) [] int {
	var output []int
	for _, n := range nums {
		if n%2 !=0 {
			output =append(output, n)
		}
	}
	return output
} 

func primeNumbers (nums [] int) [] int {
	var output []int
	for _, n := range nums {
		if isPrime(n) {
			output =append(output, n)
		}
	}
	return output
)

func oddPrimeNumbers(nums [] int) [] int {
	var output []int
	for _, n := range nums {
		if n%2 ==0 {
			output =append(output, n)
		}
	}
	return output
}
