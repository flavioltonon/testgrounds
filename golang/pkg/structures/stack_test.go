package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	input := []string{"7", "2", "1", "4", "9", "6", "0", "3", "8", "5"}
	expected := []string{"5", "8", "3", "0", "6", "9", "4", "1", "2", "7"}

	s := NewStack()

	for _, v := range input {
		s.Push(v)
	}

	assert.Equal(t, len(input), s.Len())

	for i := 0; s.Len() > 0; i++ {
		value := s.Pop()

		assert.Equal(t, expected[i], value)
	}

	assert.Equal(t, 0, s.Len())
}
