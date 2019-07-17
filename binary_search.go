package main


// BinarySearch implementation, accepts sorted array of int and target
func BinarySearch(array []int, t int) int {
	r := len(array)
	l := 0

	if r <= 1 {
		return -1
	}

	for l <= r {
		mid := (l + r) / 2

		if array[mid] < t {
			l = mid + 1
		} else if array[mid] > t {
			r = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
