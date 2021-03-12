package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	source := []int{7, 2, 1, 4, 9, 6, 0, 3, 8, 5}

	exists := LinearSearch(9, source)

	assert.Equal(t, true, exists)
}
