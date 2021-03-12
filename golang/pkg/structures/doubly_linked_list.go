package structures

// DoublyLinkedList represents a Linked List data structure, which should contain a series of
// nodes pointing to the next and the previous nodes.
type DoublyLinkedList struct {
	first  *DoublyLinkedListNode
	last   *DoublyLinkedListNode
	length int
}

// NewDoublyLinkedList creates a DoublyLinkedList and returns a pointer to it
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Len returns the size of the DoublyLinkedList
func (l *DoublyLinkedList) Len() int {
	return l.length
}

// Values returns all the values the DoublyLinkedList contains
func (l *DoublyLinkedList) Values() []string {
	values := make([]string, 0, l.length)

	current := l.first

	for current != nil {
		values = append(values, current.value)
		current = current.next
	}

	return values
}

// First returns a pointer to the first node in the DoublyLinkedList
func (l *DoublyLinkedList) First() *DoublyLinkedListNode {
	return l.first
}

// Last returns a pointer to the last node in the DoublyLinkedList
func (l *DoublyLinkedList) Last() *DoublyLinkedListNode {
	return l.last
}

// Search returns true if the input value exist in any of its nodes and returns false otherwise in
// an O(n) time complexity.
func (l *DoublyLinkedList) Search(value string) bool {
	return l.first.search(value)
}

// Prepend adds a new value to the beginning of the DoublyLinkedList with an O(1) time complexity
func (l *DoublyLinkedList) Prepend(value string) {
	new := &DoublyLinkedListNode{
		value: value,
	}

	if l.first != nil {
		new.next = l.first
		new.next.previous = new
	} else {
		l.last = new
	}

	l.first = new
	l.length++
}

// Append adds a new value to the end of the DoublyLinkedList with an O(1) time complexity
func (l *DoublyLinkedList) Append(value string) {
	new := &DoublyLinkedListNode{
		value: value,
	}

	if l.last != nil {
		l.last.next = new
		new.previous = l.last
	} else {
		l.first = new
	}

	l.last = new
	l.length++
}

// Delete removes a value from the DoublyLinkedList (if it exists), minding the order of the nodes values
func (l *DoublyLinkedList) Delete(value string) bool {
	n := l.first

	for n != nil {
		if n.value != value {
			n = n.next
			continue
		}

		if n.next == nil {
			l.last = n.previous
		} else {
			n.next.previous = n.previous
		}

		if n.previous == nil {
			l.first = n.next
		} else {
			n.previous.next = n.next
		}

		l.length--

		return true
	}

	return false
}

// DoublyLinkedListNode represents a single node of Doubly Linked List data structures, which should contain
// a value and a pointer to the next and the previous DoublyLinkedListNodes
type DoublyLinkedListNode struct {
	value    string
	next     *DoublyLinkedListNode
	previous *DoublyLinkedListNode
}

// Value returns the value stored on the current node
func (n *DoublyLinkedListNode) Value() string {
	return n.value
}

// Next returns a pointer to the node next to the current node
func (n *DoublyLinkedListNode) Next() *DoublyLinkedListNode {
	return n.next
}

// Previous returns a pointer to the node previous to the current node
func (n *DoublyLinkedListNode) Previous() *DoublyLinkedListNode {
	return n.previous
}

func (n *DoublyLinkedListNode) search(value string) bool {
	if n == nil {
		return false
	}

	if n.value != value {
		return n.next.search(value)
	}

	return true
}
