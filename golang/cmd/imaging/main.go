package main

import (
	"fmt"
	"testgrounds/imaging"
)

func main() {
	path := fmt.Sprintf("assets/images/sample.jpg")

	err := imaging.Process(path,
		imaging.Resize(1024, 768),
		imaging.Grayscale(),
		imaging.AdjustContrast(75),
		imaging.AdjustBrightness(-15),
		imaging.Sharpen(3),
	)
	if err != nil {
		panic(err)
	}
}
