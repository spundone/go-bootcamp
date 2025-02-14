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
            filterFn: evenNumbers,
            expected: []int{2, 4, 6, 8, 10},
        },
        {
            name:     "OddNumbers",
            filterFn: oddNumbers,
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
            name:     "OddMultiplesOf3",
            	filterFn: oddMultiplesOf3,
            expected: []int{15},
        },
		// {
		// 	name:     "EvenMultiplesOf5AndOddMultiplesOf3",
		// 	filterFn: filterNumbers(inputExtended, []func(int) bool{filterEvenMultiplesOf5, filterOddMultiplesOf3}),
		// 	expected: []int{10, 15},
		// },
		
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var result []int
            if tt.name == "EvenMultiplesOf5" || tt.name == "OddMultiplesOf3" {
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
