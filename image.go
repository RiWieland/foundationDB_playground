package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"log"
	"os"
)

type editableImage struct {
	draw.Image
}

type Img struct {
	size []image.Point
}

func readImg(path string) editableImage {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img_, _ := drawableRGBImage(f)
	custImg := editableImage{
		img_,
	}
	if err != nil {
		log.Fatal(err)
	}
	return custImg
}

func writeImg(path string, img draw.Image) {
	out, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	e := jpeg.Encode(out, img, nil)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

}

func addRectangle(img editableImage, coor rectCoord) draw.Image {

	myRectangle := image.Rect(coor.x0, coor.y0, coor.x1, coor.y1)

	myColor := color.RGBA{255, 0, 255, 255}

	min := myRectangle.Min
	max := myRectangle.Max

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

func drawableRGBImage(f io.Reader) (draw.Image, error) {
	img, err := jpeg.Decode(f)
	if err != nil {
		return nil, err
	}
	b := img.Bounds()
	output_rgb := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(output_rgb, output_rgb.Bounds(), img, b.Min, draw.Src)

	return output_rgb, nil
}
