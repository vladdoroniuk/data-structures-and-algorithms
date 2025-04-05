package selection_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSortInt(t *testing.T) {
	ints := []int{5, 3, 9, 1, 7}
	expected := []int{1, 3, 5, 7, 9}

	assert.Equal(t, expected, selectionSort(ints))
}

func TestSelectionSortString(t *testing.T) {
	strings := []string{"banana", "apple", "date", "cherry", "fig"}
	expected := []string{"apple", "banana", "cherry", "date", "fig"}

	assert.Equal(t, expected, selectionSort(strings))
}

func TestSelectionSortFloat(t *testing.T) {
	floats := []float64{3.3, 1.1, 5.5, 4.4, 2.2}
	expected := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	assert.Equal(t, expected, selectionSort(floats))
}

func TestSelectionSortEmpty(t *testing.T) {
	var empty []int
	var expected []int

	assert.Equal(t, expected, selectionSort(empty))
}

func TestSelectionSortSingleElement(t *testing.T) {
	single := []int{42}
	expected := []int{42}

	assert.Equal(t, expected, selectionSort(single))
}

func TestSelectionSortAlreadySorted(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	assert.Equal(t, expected, selectionSort(sorted))
}

func TestSelectionSortReverseSorted(t *testing.T) {
	reverse := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}

	assert.Equal(t, expected, selectionSort(reverse))
}

func TestSelectionSortNilSlice(t *testing.T) {
	var nilSlice []int
	assert.Nil(t, selectionSort(nilSlice))
}
