package sort

// MergeSort performs a MergeSort algorithm, sorting an slice with O(n.log(n)) time complexity
func MergeSort(input []int) {
	size := len(input)

	if size == 0 || size == 1 {
		return
	}

	// Sort left side
	left := input[:size/2]
	MergeSort(left)

	// Sort right side
	right := input[size/2:]
	MergeSort(right)

	// Merge sorted results
	copy(input, merge(left, right))
}

func merge(left []int, right []int) []int {
	size := len(left) + len(right)
	merged := make([]int, 0, size)

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	merged = append(merged, left...)
	merged = append(merged, right...)

	return merged
}
