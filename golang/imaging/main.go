package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

func main() {
	filename := "rg-verso.jpg"

	src, err := imaging.Open(fmt.Sprintf("assets/input/%s", filename))
	if err != nil {
		panic(fmt.Sprintf("failed to open image: %v", err))
	}

	w := 1024
	h := 768
	if src.Bounds().Dx() > src.Bounds().Dy() {
		h = 0
	} else {
		w = 0
	}

	t0 := time.Now()
	src = imaging.Resize(src, w, h, imaging.Lanczos)
	fmt.Println(fmt.Sprintf("Resizing took %v milliseconds", time.Since(t0).Milliseconds()))

	// Create a blurred version of the image.
	new := imaging.Grayscale(src)

	// // Create a grayscale version of the image with higher contrast and sharpness.
	new = imaging.AdjustContrast(new, 75)
	new = imaging.AdjustBrightness(new, -15)
	new = imaging.Sharpen(new, 3)

	var buf bytes.Buffer

	err = imaging.Encode(&buf, new, imaging.JPEG, imaging.JPEGQuality(100))
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(&buf)
	if err != nil {
		panic(err)
	}

	var b64 []byte
	_, err = base64.StdEncoding.Decode(b64, b)
	if err != nil {
		panic(err)
	}

	fmt.Println(b64)

	// Save the resulting image as JPEG.
	err = imaging.Save(new, fmt.Sprintf("assets/output/%s%s", uuid.New().String(), filepath.Ext(filename)))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
