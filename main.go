package main

import (
	"fmt"

	"image"
	"time"
	/*
	  "github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	  "github.com/apple/foundationdb/bindings/go/src/fdb/tuple"

	  "fmt"
	  "strconv"
	  "errors"
	  "sync"
	  "math/rand"
	*/)

// To do:
// - transfer data between routines -> channnel /chan
// - make sure every person only got one account
// - implement restrictions:
// -  - no negative vales
// - overwriting keys in key-value store OR timestamp?

func main() {

	db := kvStore{}
	db.initFdb()

	// directory raw
	//db.initDirectory("rawFiles")
	file_path := "2023-10-12T16:02:32.342Z_18:34:02.123Z_cam1.mp4"
	var f file
	fi := f.extractFileMeta(file_path)
	fmt.Println(fi)

	// TEST Func for Rectangle
	EditImg := readImg("")

	var origImg Img

	origImg.size = []image.Point{
		image.Point{10, 190},   // top-left
		image.Point{10, 240},   // bottom-left
		image.Point{1000, 200}, // bottom-right
		image.Point{1000, 150}, // top-right
	}

	out := CreateOutputFile("test.")

	img_marked := addPointVector(EditImg, origImg.size)

}

type file struct {
	startTime time.Time
	endTime   time.Time
	path      string
	object    string
	fileType  string
}

// coordinates where object is detected
type ObjectCoord struct {
}

// duration when the object is visible
type objectDuration struct {
	start time.Time
	end   time.Time
}

// draft for keyValue
type keyValue struct {
	f file
	t ObjectCoord
	d objectDuration
}
