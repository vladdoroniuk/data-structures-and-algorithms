package binarysearch

import (
	"cmp"
)

func BinarySearch[S ~[]E, E cmp.Ordered](slice S, target E) int {
	low, high := 0, len(slice)-1

	for low <= high {
		mid := low + (high-low)/2 // (high+low)/2 can cause integer overflow

		if slice[mid] < target {
			low = mid + 1 // exclude current mid index
		} else if slice[mid] > target {
			high = mid - 1 // exclude current mid index
		} else {
			return mid
		}
	}

	return -1 // not found
}
