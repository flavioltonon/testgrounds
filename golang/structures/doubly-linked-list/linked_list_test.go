package doubly_test

import (
	"testing"

	doubly "testgrounds/structures/doubly-linked-list"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	input := []string{"7", "2", "1", "4", "9", "6", "0", "3", "8", "5"}
	expected := []string{"5", "8", "3", "0", "6", "9", "4", "1", "2", "7"}

	list := doubly.NewLinkedList()

	for _, v := range input {
		list.Prepend(v)
	}

	assert.Equal(t, len(input), list.Len())

	for i, value := range list.Values() {
		assert.Equal(t, expected[i], value)
	}

	assert.Equal(t, false, list.Search("10"))
	assert.Equal(t, true, list.Search("1"))
	assert.Equal(t, true, list.Search("5"))

	assert.Equal(t, true, list.Delete("3"))
	assert.Equal(t, false, list.Search("3"))
}
