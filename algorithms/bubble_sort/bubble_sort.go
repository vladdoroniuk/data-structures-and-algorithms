package bubble_sort

import "cmp"

func bubbleSort[S ~[]E, E cmp.Ordered](slice S) S {
	if len(slice) < 2 {
		return slice
	}

	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}
