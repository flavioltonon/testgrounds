package structures

// SinglyLinkedList represents a Linked List data structure, which should contain a series of
// nodes pointing to the next node.
type SinglyLinkedList struct {
	head   *SinglyLinkedListNode
	length int
}

// NewSinglyLinkedList creates a SinglyLinkedList and returns a pointer to it
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Len returns the size of the SinglyLinkedList
func (l *SinglyLinkedList) Len() int {
	return l.length
}

// Values returns all the values the SinglyLinkedList contains
func (l *SinglyLinkedList) Values() []string {
	values := make([]string, 0, l.length)

	current := l.head

	for current != nil {
		values = append(values, current.value)
		current = current.next
	}

	return values
}

// Head returns a pointer to the first node in the SinglyLinkedList
func (l *SinglyLinkedList) Head() *SinglyLinkedListNode {
	return l.head
}

// Search returns true if the input value exist in any of its nodes and returns false otherwise in
// an O(n) time complexity.
func (l *SinglyLinkedList) Search(value string) bool {
	return l.head.search(value)
}

// Add adds a new value to the beginning of the SinglyLinkedList with an O(1) time complexity
func (l *SinglyLinkedList) Prepend(value string) {
	new := &SinglyLinkedListNode{
		value: value,
	}

	if l.head != nil {
		new.next = l.head
	}

	l.head = new
	l.length++
}

// Delete removes a value from the SinglyLinkedList (if it exists), minding the order of the nodes values
func (l *SinglyLinkedList) Delete(value string) bool {
	if l.head == nil {
		return false
	}

	if l.head.value == value {
		l.head = l.head.next
	} else if deleted := l.head.delete(value); !deleted {
		return false
	}

	l.length--

	return true
}

// SinglyLinkedListNode represents a node, which should contain a value and a pointer to a next SinglyLinkedListNode
type SinglyLinkedListNode struct {
	value string
	next  *SinglyLinkedListNode
}

// Value returns the value stored on the current node
func (n *SinglyLinkedListNode) Value() string {
	return n.value
}

// Next returns a pointer to the node next to the current node
func (n *SinglyLinkedListNode) Next() *SinglyLinkedListNode {
	return n.next
}

func (n *SinglyLinkedListNode) search(value string) bool {
	if n == nil {
		return false
	}

	if n.value != value {
		return n.next.search(value)
	}

	return true
}

func (n *SinglyLinkedListNode) delete(value string) bool {
	if n.next == nil {
		return false
	}

	if n.next.value != value {
		return n.next.delete(value)
	}

	n.next = n.next.next

	return true
}
