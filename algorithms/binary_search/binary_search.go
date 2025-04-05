package binary_search

import (
	"cmp"
)

func binarySearch[S ~[]E, E cmp.Ordered](slice S, target E) int {
	low, high := 0, len(slice)-1

	for low <= high {
		mid := low + (high-low)/2 // (high+low)/2 can cause integer overflow

		if target > slice[mid] {
			low = mid + 1
		} else if target < slice[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
