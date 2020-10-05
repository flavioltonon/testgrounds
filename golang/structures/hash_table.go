package structures

import (
	"unicode"
)

// HashTableSize defines the size of the base array of HashTables
const HashTableSize = 26

type HashTable struct {
	Buckets [HashTableSize]*LinkedList
}

// NewHashTable creates a HashTable and returns a pointer to it
func NewHashTable() *HashTable {
	return &HashTable{}
}

// Size returns the number of elements the entire HashTable contains, considering all its buckets
func (h HashTable) Size() int {
	var size int

	for _, bucket := range h.Buckets {
		if bucket == nil {
			continue
		}

		size += bucket.Len()
	}

	return size
}

// Search returns true if the input value exist in any of its nodes and returns false otherwise in
// an O(1 + n/k) time complexity, where n represents the size of the bucket and k is the HashTableSize.
// When k is large enough (such as Go native maps - 1 key -> 1 value), it drops down to O(1). When there is a single
// bucket available, it becomes O(n).
func (h *HashTable) Search(value string) bool {
	index := h.hash(value)

	if h.Buckets[index] == nil {
		return false
	}

	return h.Buckets[index].Search(value)
}

// Insert adds a new value to the HashTable, minding the order of the nodes values in the target Bucket
func (h *HashTable) Insert(value string) {
	index := h.hash(value)

	if h.Buckets[index] == nil {
		h.Buckets[index] = NewLinkedList()
	}

	h.Buckets[index].Insert(value)
}

// Delete removes a value from the HashTable (if it exists), minding the order of the nodes values in the target Bucket
func (h *HashTable) Delete(value string) bool {
	index := h.hash(value)

	if h.Buckets[index] == nil {
		return false
	}

	return h.Buckets[index].Delete(value)
}

// hash represents the hash function which is used to map input values to an specific
// Bucket index on the HashTable
func (h HashTable) hash(value string) int {
	s0 := unicode.ToLower(rune(value[0]))
	return int(s0) % HashTableSize
}
