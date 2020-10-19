package structures_test

import (
	"testing"

	"testgrounds/structures"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	input := []string{"7", "2", "1", "4", "9", "6", "0", "3", "8", "5"}
	expected := []string{"5", "8", "3", "0", "6", "9", "4", "1", "2", "7"}

	linkedList := structures.NewLinkedList()

	for _, v := range input {
		linkedList.Add(v)
	}

	assert.Equal(t, len(input), linkedList.Len())

	for i, value := range linkedList.Values() {
		assert.Equal(t, expected[i], value)
	}

	assert.Equal(t, false, linkedList.Search("10"))
	assert.Equal(t, true, linkedList.Search("1"))
	assert.Equal(t, true, linkedList.Search("5"))
}
