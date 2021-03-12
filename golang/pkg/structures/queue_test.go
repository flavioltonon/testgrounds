package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	input := []string{"7", "2", "1", "4", "9", "6", "0", "3", "8", "5"}
	expected := []string{"7", "2", "1", "4", "9", "6", "0", "3", "8", "5"}

	s := NewQueue()

	for _, v := range input {
		s.Enqueue(v)
	}

	assert.Equal(t, len(input), s.Len())

	for i := 0; s.Len() > 0; i++ {
		value := s.Dequeue()

		assert.Equal(t, expected[i], value)
	}

	assert.Equal(t, 0, s.Len())
}
