package main

import (
	"fmt"
	"image/color"

	"golang.org/x/image/colornames"
	"gopkg.in/go-playground/colors.v1"
)

var svg color.Palette

func init() {
	svg = make(color.Palette, 0)

	for _, c := range colornames.Map {
		svg = append(svg, c)
	}
}

func main() {
	h := "#9ACD32"

	c, err := colors.Parse(h)
	if err != nil {
		panic(err)
	}

	tmp := c.ToRGBA()

	rgba := color.RGBA{
		R: tmp.R,
		G: tmp.G,
		B: tmp.B,
		A: uint8(tmp.A),
	}

	fmt.Println(rgba.RGBA())

	match := svg.Convert(rgba)

	fmt.Println(match.RGBA())

	fmt.Println(match == colornames.Yellowgreen)

	// Display all colors
	displayColors()
}

func displayColors() {
	for _, c := range colornames.Map {
		// RGBA to YCbCR
		tmp := color.YCbCrModel.Convert(c).(color.YCbCr)

		// YCbCr to RGB
		r, g, b := color.YCbCrToRGB(tmp.Y, tmp.Cb, tmp.Cr)

		fmt.Println(r, g, b)
	}
}
