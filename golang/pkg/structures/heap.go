package structures

// Heap represents a heap data structure, which should contain a series of
// nodes pointing to other two nodes: left and right.
type Heap struct {
	root *HeapNode
	size int
}

// NewHeap creates a Heap and returns a pointer to it
func NewHeap() *Heap {
	return &Heap{}
}

// Root returns a pointer to the root node in the Heap
func (h Heap) Root() *HeapNode {
	return h.root
}

// Size returns the number of nodes in the Heap
func (h Heap) Size() int {
	return h.size
}

// Insert adds a value to the Heap
func (h *Heap) Insert(value string) {
	currentNode := h.root

	for currentNode != nil {
		if value < currentNode.value {
			currentNode = currentNode.left
		} else if value > currentNode.value {
			currentNode = currentNode.right
		} else {
			return
		}
	}

	currentNode = &HeapNode{value: value}
}

// Delete removes a value from the Heap (if it exists), minding the order of the nodes values
func (h *Heap) Delete(value string) bool {
	return false
}

// Search returns true if the input value exist in any of its nodes and returns false otherwise in
// an O(log(n)) time complexity.
func (h Heap) Search(value string) bool {
	return false
}

// HeapNode represents a single node of heap data structures, which should contain
// a value and pointers to its children
type HeapNode struct {
	value string
	left  *HeapNode
	right *HeapNode
}

// Value returns the value stored on the current node
func (n HeapNode) Value() string {
	return n.value
}

// Left returns a pointer to the node's left child
func (n HeapNode) Left() *HeapNode {
	return n.left
}

// Right returns a pointer to the node's right child
func (n HeapNode) Right() *HeapNode {
	return n.right
}
