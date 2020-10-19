package stack

// Stack represents a stack data structure
type Stack struct {
	elements []string
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]string, 0),
	}
}

func (s *Stack) WithCapacity(cap int) *Stack {
	return &Stack{
		elements: make([]string, 0, cap),
	}
}

func (s *Stack) Len() int {
	return len(s.elements)
}

func (s *Stack) Push(value string) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() string {
	n := len(s.elements) - 1

	popped := s.elements[n]

	s.elements = s.elements[:n]

	return popped
}
