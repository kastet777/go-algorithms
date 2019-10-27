package main

import "testing"

func TestBinatySearch(t *testing.T) {
	arr := []int{1, 3, 5, 6, 7, 8, 8, 12}

	targetValue := 6
	expectedIndex := 3

	result := BinarySearch(arr, targetValue)

	if result != expectedIndex {
		t.Error("Result doesn't match expected index")
	}
}
