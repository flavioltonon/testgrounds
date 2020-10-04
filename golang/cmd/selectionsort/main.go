package main

import (
	"fmt"
	"testgrounds/selectionsort"
)

func main() {
	input := []int{7, 2, 1, 4, 9, 0, 3, 5}

	selectionsort.Sort(input)

	fmt.Println(input)
}
