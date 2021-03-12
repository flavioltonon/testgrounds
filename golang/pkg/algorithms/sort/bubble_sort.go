package sort

// BubbleSort performs a BubbleSort algorithm, sorting an slice with O(n^2) time complexity
func BubbleSort(input []int) {
	if len(input) < 2 {
		// Nothing to sort
		return
	}

	for i := 0; i < len(input); i++ {
		// Find the largest numbers in the slice and carry them one by one to the position they should be
		// in order to sort the input values
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				// Swap
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
}
