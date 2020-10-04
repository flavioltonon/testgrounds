package mergesort_test

import (
	"fmt"
	"testing"

	"testgrounds/mergesort"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	input := []int{7, 2, 1, 4, 9, 6, 0, 3, 8, 5}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	mergesort.Sort(input)

	fmt.Printf("Result: %v\n", input)

	for i, v := range input {
		assert.Equal(t, expected[i], v)
	}
}
