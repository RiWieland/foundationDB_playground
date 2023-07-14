package main


import (
	"image/draw""
)

type CustomImage struct {
	draw.Image
}

func addRectangle(img CustomImage, rect image.Rectangle) draw.Image {
	myColor := color.RGBA{255, 0, 255, 255}

	min := rect.Min
	max := rect.Max

	for i := min.X; i < max.X; i++ {
		img.Set(i, min.Y, myColor)
		img.Set(i, max.Y, myColor)
	}

	for i := min.Y; i <= max.Y; i++ {
		img.Set(min.X, i, myColor)
		img.Set(max.X, i, myColor)
	}
	return img
}
