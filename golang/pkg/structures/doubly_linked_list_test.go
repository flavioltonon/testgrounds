package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList(t *testing.T) {
	input := []string{"7", "2", "1", "4", "9", "6", "0", "3", "8", "5"}

	t.Run("Prepend", func(t *testing.T) {
		expected := []string{"5", "8", "3", "0", "6", "9", "4", "1", "2", "7"}

		list := NewDoublyLinkedList()

		for _, v := range input {
			list.Prepend(v)
		}

		assert.Equal(t, len(input), list.Len())

		for i, value := range list.Values() {
			assert.Equal(t, expected[i], value)
		}

		assert.Equal(t, false, list.Search("10"))
		assert.Equal(t, true, list.Search("1"))
		assert.Equal(t, true, list.Delete("5"))
		assert.Equal(t, false, list.Search("5"))
		assert.Equal(t, true, list.Delete("7"))
		assert.Equal(t, false, list.Search("7"))
	})

	t.Run("Append", func(t *testing.T) {
		expected := []string{"7", "2", "1", "4", "9", "6", "0", "3", "8", "5"}

		list := NewDoublyLinkedList()

		for _, v := range input {
			list.Append(v)
		}

		assert.Equal(t, len(input), list.Len())

		for i, value := range list.Values() {
			assert.Equal(t, expected[i], value)
		}

		assert.Equal(t, false, list.Search("10"))
		assert.Equal(t, true, list.Search("1"))
		assert.Equal(t, true, list.Delete("5"))
		assert.Equal(t, false, list.Search("5"))
		assert.Equal(t, true, list.Delete("7"))
		assert.Equal(t, false, list.Search("7"))
	})
}
