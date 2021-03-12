package search

import (
	"testing"

	"testgrounds/pkg/algorithms/sort"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	source := []int{7, 2, 1, 4, 9, 6, 0, 3, 8, 5}

	sort.MergeSort(source)

	assert.Equal(t, true, BinarySearch(7, source))
}
