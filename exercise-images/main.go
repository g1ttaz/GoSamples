package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image is a description of an image
type Image struct {
	Width, Height int
	f             int
}

// ColorModel returns the color model of an image
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the size of an image
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

// At returns the color at a certain point in the image
func (i Image) At(x, y int) color.Color {
	var val int
	switch i.f {
	case 1:
		val = (x + y) / 2
	case 2:
		val = x * y
	case 3:
		val = x ^ y
	}

	return color.RGBA{uint8(val), uint8(val), 255, 255}
}

func main() {
	m := Image{100, 100, 1}
	pic.ShowImage(m)
}
