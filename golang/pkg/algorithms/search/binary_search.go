package search

// BinarySearch runs a binary search trying to find a value on an input array with a O(n*log(n)) time complexity,
// where n is the size of the source array. This kind of search may not work if the source is not sorted.
func BinarySearch(value int, source []int) bool {
	size := len(source)

	if size == 0 {
		return false
	}

	if value == source[size/2] {
		return true
	}

	if value < source[size/2] {
		return BinarySearch(value, source[:size/2])
	}

	if value > source[size/2] {
		return BinarySearch(value, source[1+size/2:])
	}

	return false
}
