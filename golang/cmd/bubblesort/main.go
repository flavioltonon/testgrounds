package main

import (
	"fmt"
	"testgrounds/bubblesort"
)

func main() {
	input := []int{7, 2, 1, 4, 9, 0, 3, 5}

	bubblesort.Sort(input)

	fmt.Println(input)
}
