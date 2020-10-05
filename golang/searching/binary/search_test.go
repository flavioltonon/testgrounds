package binary_test

import (
	"testing"

	"testgrounds/searching/binary"
	"testgrounds/sorting/merge"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	source := []int{7, 2, 1, 4, 9, 6, 0, 3, 8, 5}

	merge.Sort(source)

	assert.Equal(t, true, binary.Search(7, source))
}
