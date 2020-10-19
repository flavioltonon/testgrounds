package stack

// Stack represents a stack data structure, which is characterized by allowing the
// LIFO (Last In First Out) handling elements in a list
type Stack struct {
	elements []string
}

// NewStack creates a new Stack and returns a pointer to it
func NewStack() *Stack {
	return &Stack{
		elements: make([]string, 0),
	}
}

// WithCapacity allows customization of the Stack's underlying slice size
func (s *Stack) WithCapacity(cap int) *Stack {
	return &Stack{
		elements: make([]string, 0, cap),
	}
}

// Len returns the number of elements currently in the Stack
func (s *Stack) Len() int {
	return len(s.elements)
}

// Push adds a new element to the end of the Stack
func (s *Stack) Push(value string) {
	s.elements = append(s.elements, value)
}

// Pop removes the last element in the Stack and returns its value
func (s *Stack) Pop() string {
	n := len(s.elements) - 1

	popped := s.elements[n]

	s.elements = s.elements[:n]

	return popped
}
