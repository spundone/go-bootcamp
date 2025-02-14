package main

import (
    "reflect"
    "testing"
)

func TestFilterEvenNumbers(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    expected := []int{2, 4, 6, 8, 10}
    result := filterEvenNumbers(input)
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, but got %v", expected, result)
    }
}

func TestFilterOddNumbers(t *testing.T) {
    input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    expected := []int{1, 3, 5, 7, 9}
    result := filterOddNumbers(input)
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, but got %v", expected, result)
    }
}
