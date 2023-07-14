package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
)

type customImage struct {
	draw.Image
}

func readImg(path string) customImage {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img_, _ := drawableRGBImage(f)
	custImg := customImage{
		img_,
	}
	if err != nil {
		log.Fatal(err)
	}
	return custImg
}

func addRectangle(img customImage, rect image.Rectangle) draw.Image {
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
