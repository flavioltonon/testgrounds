package main

import (
	"fmt"
	"testgrounds/mergesort"
)

func main() {
	input := []int{7, 2, 1, 4, 9, 0, 3, 5}

	mergesort.Sort(input)

	fmt.Println(input)
}
