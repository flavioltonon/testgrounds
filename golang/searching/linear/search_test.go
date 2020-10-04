package linear_test

import (
	"testing"

	"testgrounds/searching/linear"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	source := []int{7, 2, 1, 4, 9, 6, 0, 3, 8, 5}

	exists := linear.Search(9, source)

	assert.Equal(t, true, exists)
}
