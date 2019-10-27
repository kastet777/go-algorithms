package main

import (
	"testing"
	"reflect"
)

func TestQuickSort(t *testing.T) {
	arr := []int{33, 42, 1, 2, 4, 5, 21, 55, 2, 19}
	expectedArr := []int{1, 2, 2, 4, 5 ,19, 21, 33, 42, 55}

	sortedArr := QuickSort(arr)

	if !reflect.DeepEqual(sortedArr, expectedArr) {
		t.Error("Result doesn't equal expected result")
	}
}
