package main

import (
	"fmt"
	"testgrounds/sorting/selection"
)

func main() {
	input := []int{7, 2, 1, 4, 9, 0, 3, 5}

	selection.Sort(input)

	fmt.Println(input)
}
