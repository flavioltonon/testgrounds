package singly_linked_list

// LinkedList represents a Linked List data structure, which should contain a series of
// nodes pointing to the next node.
type LinkedList struct {
	head   *Node
	length int
}

// NewLinkedList creates a LinkedList and returns a pointer to it
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Len returns the size of the LinkedList
func (l *LinkedList) Len() int {
	return l.length
}

// Values returns all the values the LinkedList contains
func (l *LinkedList) Values() []string {
	values := make([]string, 0, l.length)

	current := l.head

	for current != nil {
		values = append(values, current.value)
		current = current.next
	}

	return values
}

// Head returns a pointer to the first node in the LinkedList
func (l *LinkedList) Head() *Node {
	return l.head
}

// Search returns true if the input value exist in any of its nodes and returns false otherwise in
// an O(n) time complexity.
func (l *LinkedList) Search(value string) bool {
	return l.head.search(value)
}

// Add adds a new value to the beginning of the LinkedList with an O(1) time complexity
func (l *LinkedList) Prepend(value string) {
	new := &Node{
		value: value,
	}

	if l.head != nil {
		new.next = l.head
	}

	l.head = new
	l.length++
}

// Delete removes a value from the LinkedList (if it exists), minding the order of the nodes values
func (l *LinkedList) Delete(value string) bool {
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

// Node represents a single node of Linked List data structures, which should contain
// a value and a pointer to the next Node
type Node struct {
	value string
	next  *Node
}

// Value returns the value stored on the current node
func (n *Node) Value() string {
	return n.value
}

// Next returns a pointer to the node next to the current node
func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) search(value string) bool {
	if n == nil {
		return false
	}

	if n.value != value {
		return n.next.search(value)
	}

	return true
}

func (n *Node) delete(value string) bool {
	if n.next == nil {
		return false
	}

	if n.next.value != value {
		return n.next.delete(value)
	}

	n.next = n.next.next

	return true
}
