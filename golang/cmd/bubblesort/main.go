package main

import (
	"fmt"
	"testgrounds/sorting/bubble"
)

func main() {
	input := []int{7, 2, 1, 4, 9, 0, 3, 5}

	bubble.Sort(input)

	fmt.Println(input)
}
