package main

import (
	"fmt"
)

func main() {

	arr := []int{1,3,5,6,7,8,8,12}

	target := 7
	binarySearchResult := BinarySearch(arr, target)

	fmt.Printf("BinarySearch result: %v for given array %v with target value %v \n\n", binarySearchResult, arr, target)


	arr2 := []int{33,42,1,2,4,5,21,55,2,19}

	quickSortResult := QuickSort(arr2)

	fmt.Printf("QuickSort result: %v for given array %v \n\n", quickSortResult, arr2)
}