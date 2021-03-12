package structures

// Queue represents a queue data structure, which is characterized by allowing the
// FIFO (First In First Out) handling elements in a list
type Queue struct {
	elements []string
}

// NewQueue creates a new Queue and returns a pointer to it
func NewQueue() *Queue {
	return &Queue{
		elements: make([]string, 0),
	}
}

// WithCapacity allows customization of the Queue's underlying slice size
func (s *Queue) WithCapacity(cap int) *Queue {
	return &Queue{
		elements: make([]string, 0, cap),
	}
}

// Len returns the number of elements currently in the Queue
func (s *Queue) Len() int {
	return len(s.elements)
}

// Enqueue adds a new element to the end of the Queue
func (s *Queue) Enqueue(value string) {
	s.elements = append(s.elements, value)
}

// Dequeue removes the first element in the Queue and returns its value
func (s *Queue) Dequeue() string {
	popped := s.elements[0]

	s.elements = s.elements[1:]

	return popped
}
