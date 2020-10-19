package doubly

// DoublyLinkedList represents a Linked List data structure, which should contain a series of
// nodes pointing to the next node.
type DoublyLinkedList struct {
	head   *Node
	length int
}

// NewList creates a DoublyLinkedList and returns a pointer to it
func NewLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Len returns the size of the DoublyLinkedList
func (l *DoublyLinkedList) Len() int {
	return l.length
}

// Values returns all the values the DoublyLinkedList contains
func (l *DoublyLinkedList) Values() []string {
	values := make([]string, 0, l.length)

	current := l.head

	for current != nil {
		values = append(values, current.value)
		current = current.next
	}

	return values
}

// Head returns a pointer to the first node in the DoublyLinkedList
func (l *DoublyLinkedList) Head() *Node {
	return l.head
}

// Search returns true if the input value exist in any of its nodes and returns false otherwise in
// an O(n) time complexity.
func (l *DoublyLinkedList) Search(value string) bool {
	return l.head.search(value)
}

// Add adds a new value to the beginning of the DoublyLinkedList with an O(1) time complexity
func (l *DoublyLinkedList) Prepend(value string) {
	new := &Node{
		value: value,
	}

	if l.head != nil {
		new.next = l.head
		new.next.previous = new
	}

	l.head = new
	l.length++
}

// Delete removes a value from the DoublyLinkedList (if it exists), minding the order of the nodes values
func (l *DoublyLinkedList) Delete(value string) bool {
	deleted := l.head.delete(value)

	if !deleted {
		return false
	}

	l.length--

	return true
}

// Node represents a single node of Doubly Linked List data structures, which should contain
// a value and a pointer to the next and the previous Nodes
type Node struct {
	value    string
	next     *Node
	previous *Node
}

// Value returns the value stored on the current node
func (n *Node) Value() string {
	return n.value
}

// Next returns a pointer to the node previous to the current node
func (n *Node) Previous() *Node {
	return n.previous
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
	if n == nil {
		return false
	}

	if n.value != value {
		return n.next.delete(value)
	}

	n.next.previous = n.previous
	n.previous.next = n.next

	return true
}
