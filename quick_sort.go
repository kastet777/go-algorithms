package main


// QuickSort algorithm implementation, takes array of int
func QuickSort(arr []int) []int {
	arrayLen := len(arr)

	if arrayLen <= 1 {
		return arr
	}

	pivot := arr[arrayLen - 1]

	left := []int{}
	right := []int{}

	for i := 0; i < arrayLen - 1; i++ {
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else if arr[i] > pivot {
			right = append(right, arr[i])
		}
	}

	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
