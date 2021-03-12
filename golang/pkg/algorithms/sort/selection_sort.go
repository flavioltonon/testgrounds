package sort

// SelectionSort performs a SelectionSort algorithm, sorting an slice with O(n^2) time complexity
func SelectionSort(input []int) {
	if len(input) < 2 {
		// Nothing to sort
		return
	}

	for i := 0; i < len(input); i++ {
		// Find the smallest numbers in the slice one by one and carry them to their sorted positions
		smallestValue := input[i]
		smallestValueIndex := i

		for j := i; j < len(input); j++ {
			if input[j] < smallestValue {
				smallestValue = input[j]
				smallestValueIndex = j
			}
		}

		input[i], input[smallestValueIndex] = input[smallestValueIndex], input[i]
	}
}
