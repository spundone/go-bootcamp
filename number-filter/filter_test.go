package main

import (
    "reflect"
    "testing"
)

func TestFilter(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    tests := []struct {
        name     string
        filterFn func([]int) []int
        expected []int
    }{
        {
            name:     "EvenNumbers",
            filterFn: filterEvenNumbers,
            expected: []int{2, 4, 6, 8, 10},
        },
        {
            name:     "OddNumbers",
            filterFn: filterOddNumbers,
            expected: []int{1, 3, 5, 7, 9},
        },
        {
            name:     "PrimeNumbers",
            filterFn: filterPrimeNumbers,
            expected: []int{2, 3, 5, 7},
        },
        {
            name:     "OddPrimeNumbers",
            filterFn: filterOddPrimeNumbers,
            expected: []int{3, 5, 7},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := tt.filterFn(input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("Expected %v, but got %v", tt.expected, result)
            }
        })
    }
}
