package bubble_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	ints := []int{5, 3, 9, 1, 7}
	expected := []int{1, 3, 5, 7, 9}

	assert.Equal(t, expected, bubbleSort(ints))
}

func TestBubbleSortString(t *testing.T) {
	strings := []string{"banana", "apple", "date", "cherry", "fig"}
	expected := []string{"apple", "banana", "cherry", "date", "fig"}

	assert.Equal(t, expected, bubbleSort(strings))
}

func TestBubbleSortFloat(t *testing.T) {
	floats := []float64{3.3, 1.1, 5.5, 4.4, 2.2}
	expected := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	assert.Equal(t, expected, bubbleSort(floats))
}

func TestBubbleSortEmpty(t *testing.T) {
	var empty []int
	var expected []int

	assert.Equal(t, expected, bubbleSort(empty))
}

func TestBubbleSortSingleElement(t *testing.T) {
	single := []int{42}
	expected := []int{42}

	assert.Equal(t, expected, bubbleSort(single))
}

func TestBubbleSortAlreadySorted(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	assert.Equal(t, expected, bubbleSort(sorted))
}

func TestBubbleSortReverseSorted(t *testing.T) {
	reverse := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}

	assert.Equal(t, expected, bubbleSort(reverse))
}

func TestBubbleSortNilSlice(t *testing.T) {
	var nilSlice []int
	assert.Nil(t, bubbleSort(nilSlice))
}
