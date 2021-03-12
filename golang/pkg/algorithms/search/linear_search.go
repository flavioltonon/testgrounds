package search

// LinearSearch runs a linear search trying to find a value on an input array with a O(n) time complexity,
// where n is the size of the source array.
func LinearSearch(value int, source []int) bool {
	for _, v := range source {
		if v == value {
			return true
		}
	}

	return false
}
