package structures

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
func (l LinkedList) Len() int {
	return l.length
}

// Elements returns all the elements the LinkedList contains
func (l LinkedList) Elements() []string {
	elems := make([]string, 0, l.length)

	current := l.head

	for current != nil {
		elems = append(elems, current.value)
		current = current.next
	}

	return elems
}

// Head returns a pointer to the first node in the LinkedList
func (l LinkedList) Head() *Node {
	return l.head
}

// Search returns true if the input value exist in any of its nodes and returns false otherwise in
// an O(n) time complexity.
func (l LinkedList) Search(value string) bool {
	return l.head.search(value)
}

// Insert adds a new value to the LinkedList, minding the order of the nodes values
func (l *LinkedList) Insert(value string) {
	defer func() { l.length++ }()

	new := &Node{
		value: value,
	}

	if l.head == nil {
		l.head = new
	} else if new.value < l.head.value {
		new.next = l.head
		l.head = new
	} else {
		l.head.insert(new)
	}
}

// Delete removes a value from the LinkedList (if it exists), minding the order of the nodes values
func (l *LinkedList) Delete(value string) (deleted bool) {
	defer func() {
		if deleted {
			l.length--
		}
	}()

	if l == nil {
		return false
	}

	if l.head == nil {
		return false
	}

	if l.head.value == value {
		l.head = l.head.next
		return true
	}

	return l.head.delete(value)
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

func (n *Node) insert(new *Node) {
	if n.next == nil {
		n.next = new
		return
	}

	if new.value < n.next.value {
		new.next = n.next
		n.next = new
		return
	}

	n.next.insert(new)
}

func (n *Node) delete(value string) bool {
	if n.next == nil {
		return false
	}

	if n.next.value == value {
		n.next = n.next.next
		return true
	}

	return n.next.delete(value)
}
