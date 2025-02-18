package main

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputExtended := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	tests := []struct {
		name     string
		filterFn func([]int) []int
		expected []int
	}{
		{
			name:     "EvenNumbers",
			filterFn: filterEven,
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "OddNumbers",
			filterFn: filterOdd,
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "PrimeNumbers",
			filterFn: primeNumbers,
			expected: []int{2, 3, 5, 7},
		},
		{
			name:     "OddPrimeNumbers",
			filterFn: oddPrimeNumbers,
			expected: []int{3, 5, 7},
		},
		{
			name:     "EvenMultiplesOf5",
			filterFn: evenMultiplesOf5,
			expected: []int{10, 20},
		},
		{
			name:     "OddMultiplesOf3GreaterThan10",
			filterFn: oddMultiplesOf3GreaterThan10,
			expected: []int{15},
		},
		{
			name: "AndNumber1",
			// odd, greater than 5, multiple of 3
			filterFn: func(nums []int) []int { return filtersAll(nums, odd, greaterThanN(5), multiplesOf(3)) },
			expected: []int{9, 15},
		},
		{
			name: "AndNumber2",
			// even, less than 15, multiple of 3
			filterFn: func(nums []int) []int { return filtersAll(nums, even, lessThanN(15), multiplesOf(3)) },
			expected: []int{6, 12},
		},
		{
			name: "AndNumber3",
			// prime, greater than 10, less than 15
			filterFn: func(nums []int) []int { return filtersAll(nums, prime, greaterThanN(10), lessThanN(15)) },
			expected: []int{11, 13},
		},
		{
			name:     "OrNumber1",
            // greater than 15, multiple of 5
			filterFn: func(nums []int) []int { return filterAny(nums, greaterThanN(15), multiplesOf(5)) }, //unsure about this one
			expected: []int{5, 10, 15, 16, 17, 18, 19, 20},
		},
        {
            name: "OrNumber2",
            //less than 6, multiple of 3
            filterFn: func(nums []int) []int { return filterAny(nums, lessThanN(6), multiplesOf(3)) },
            expected: []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18},
        },
        {
            name: "OrNumber3",
            // greater than 10, less than 15
            filterFn: func(nums []int) []int { return filterAny(nums, greaterThanN(10), lessThanN(15)) },
            expected: []int{1, 2, 3 ,4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
        },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result []int
			if tt.name == "EvenMultiplesOf5" || tt.name == "OddMultiplesOf3GreaterThan10" || tt.name == "AndNumber1" || tt.name == "AndNumber2" || tt.name == "AndNumber3" || tt.name == "OrNumber1" || tt.name == "OrNumber2" || tt.name == "OrNumber3" {
				result = tt.filterFn(inputExtended)
			} else {
				result = tt.filterFn(input)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
