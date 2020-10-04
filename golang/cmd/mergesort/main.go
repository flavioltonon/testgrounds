package main

import (
	"fmt"
	"testgrounds/sorting/merge"
)

func main() {
	input := []int{7, 2, 1, 4, 9, 0, 3, 5}

	merge.Sort(input)

	fmt.Println(input)
}
