package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct{}

func (I Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (I Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 200)
}

func (I Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y), uint8(x * y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
