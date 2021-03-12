package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	input := []int{7, 2, 1, 4, 9, 6, 0, 3, 8, 5}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	BubbleSort(input)

	for i, v := range input {
		assert.Equal(t, expected[i], v)
	}
}
