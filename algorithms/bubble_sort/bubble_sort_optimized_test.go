package bubble_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortOptimizedInt(t *testing.T) {
	ints := []int{5, 3, 9, 1, 7}
	expected := []int{1, 3, 5, 7, 9}

	assert.Equal(t, expected, bubbleSortOptimized(ints))
}

func TestBubbleSortOptimizedString(t *testing.T) {
	strings := []string{"banana", "apple", "date", "cherry", "fig"}
	expected := []string{"apple", "banana", "cherry", "date", "fig"}

	assert.Equal(t, expected, bubbleSortOptimized(strings))
}

func TestBubbleSortOptimizedFloat(t *testing.T) {
	floats := []float64{3.3, 1.1, 5.5, 4.4, 2.2}
	expected := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	assert.Equal(t, expected, bubbleSortOptimized(floats))
}

func TestBubbleSortOptimizedEmpty(t *testing.T) {
	var empty []int
	var expected []int

	assert.Equal(t, expected, bubbleSortOptimized(empty))
}

func TestBubbleSortOptimizedSingleElement(t *testing.T) {
	single := []int{42}
	expected := []int{42}

	assert.Equal(t, expected, bubbleSortOptimized(single))
}

func TestBubbleSortOptimizedAlreadySorted(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	assert.Equal(t, expected, bubbleSortOptimized(sorted))
}

func TestBubbleSortOptimizedReverseSorted(t *testing.T) {
	reverse := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}

	assert.Equal(t, expected, bubbleSortOptimized(reverse))
}

func TestBubbleSortOptimizedNilSlice(t *testing.T) {
	var nilSlice []int
	assert.Nil(t, bubbleSortOptimized(nilSlice))
}
