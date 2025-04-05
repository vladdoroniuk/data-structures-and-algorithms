package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchInt(t *testing.T) {
	ints := []int{1, 3, 5, 7, 9, 11}

	assert.Equal(t, 2, binarySearch(ints, 5))
}

func TestBinarySearchString(t *testing.T) {
	strings := []string{"apple", "banana", "cherry", "date", "fig", "grape"}

	assert.Equal(t, 5, binarySearch(strings, "grape"))
}

func TestBinarySearchFloat(t *testing.T) {
	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}

	assert.Equal(t, 3, binarySearch(floats, 4.4))
}
