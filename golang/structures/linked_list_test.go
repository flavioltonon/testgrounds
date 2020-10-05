package structures_test

import (
	"testing"

	"testgrounds/structures"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	input := []string{"7", "2", "1", "4", "9", "6", "0", "3", "8", "5"}
	expected := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	linkedList := structures.NewLinkedList()

	for _, v := range input {
		linkedList.Insert(v)
	}

	assert.Equal(t, len(input), linkedList.Len())

	current := linkedList.Head()

	for i := 0; i < linkedList.Len(); i++ {
		assert.Equal(t, expected[i], current.Value())
		current = current.Next()
	}

	assert.Equal(t, false, linkedList.Search("10"))
	assert.Equal(t, true, linkedList.Search("1"))
	assert.Equal(t, true, linkedList.Search("5"))
}
