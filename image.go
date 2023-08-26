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
	"time"
)

type editableImage struct {
	draw.Image
}

// imageStruct with metadata and related rect
type img struct {
	path     string
	fileType string
	time     time.Time
	rect     rectCoord
}

func readOs(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return f
}

func exportEditImage(path string) editableImage {
	f := readOs(path)
	img_, err := drawableRGBImage(f)
	if err != nil {
		log.Fatal(err)
	}
	custImg := editableImage{
		img_,
	}

	return custImg
}

// function exporting Color space from image
// to Add: direction of exporting X and Y space
// should be a method for struct?
func (img imgColor) exportImageColor(path string) imgColor {
	var r []uint8
	var g []uint8
	var b []uint8
	var a []uint8

	f := readOs(path)

	imgTemp, _, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}

	size := imgTemp.Bounds().Size()

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := imgTemp.At(x, y)
			col := color.RGBAModel.Convert(pixel).(color.RGBA)

			r = append(r, col.R)
			g = append(g, col.G)
			b = append(b, col.B)
			a = append(a, col.A)
		}
	}

	img.red = r
	img.green = g
	img.blue = b
	img.alpha = a

	return img
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

	myRectangle := image.Rect(int(coor.x0), int(coor.y0), int(coor.x1), int(coor.y1))

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
