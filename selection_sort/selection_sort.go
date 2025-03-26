package selection_sort

import "cmp"

func SelectionSort[S ~[]E, E cmp.Ordered](slice S) S {
	if len(slice) < 2 {
		return slice
	}

	for i := 0; i < len(slice)-1; i++ {
		minIdx := i

		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}

		if minIdx != i {
			slice[i], slice[minIdx] = slice[minIdx], slice[i]
		}
	}

	return slice
}
