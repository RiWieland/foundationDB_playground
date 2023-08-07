package main

import (
	"fmt"
	"log"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
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
// - query function: add methods: all and index
var imgSub subspace.Subspace
var rectSub subspace.Subspace

//var processedSub subspace.Subspace

func main() {

	var f imgMeta

	//file_path := "2023-10-12T16:02:32.342Z_18:34:02.123Z_cam1.mp4"
	file_path := "2023-10-12T18:34:00.000Z_18:34:02.123Z_cam1.mp4"
	imgMeta := f.extractImgMeta(file_path)
	fmt.Println(imgMeta.time.String())

	coor := rectCoord{
		1,
		0,
		260,
		1100,
		120,
	}

	coorN := rectCoord{
		2,
		20,
		300,
		2000,
		120,
	}

	// Put this in place for frame manipulation:

	/*
		// Image Manipulation, External Model:
		EditImg := exportEditImage("test.jpg")

		coor := rectCoord{
			0,
			260,
			1100,
			120,
		}
		img_marked := addRectangle(EditImg, coor)
		writeImg("out_rect.jpg", img_marked)


	*/
	fdbInst := kvStore{
		instance: initFdb(),
	}
	//db := initFdb()

	// add meta to file subDirectory:
	fileDir, err := directory.CreateOrOpen(fdbInst.instance, []string{"fileDir"}, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Prefex in Subs: first Element of Tuple of the key
	imgSub = fileDir.Sub("img")
	rectSub = fileDir.Sub("rect")

	// Data Model:
	// - key for the rectangle will be the coordinates
	// - no value needed

	seconds := time.Duration(10) * time.Second

	key, _ := fdbInst.writeRect(coorN)
	keyImg, _ := fdbInst.writeImgWithCoor(imgMeta, seconds, coor)
	fmt.Println(key)
	fmt.Println(keyImg)

	//t, _ := fdbInst.queryRectSub()
	i, _ := fdbInst.queryImgSub()
	fmt.Println(i)

}

// Metadata for image including rectangle coordinates
type imgMeta struct {
	path     string
	fileType string
	time     time.Time
	rect     string
}

// Image Color space
type imgColor struct {
	red   []uint8
	blue  []uint8
	green []uint8
	alpha []uint8
}

// coordinates where object is marked
type rectCoord struct {
	idx int
	x0  int
	y0  int
	x1  int
	y1  int
}

// duration when the object is visible
type objectDuration struct {
	start time.Time
	end   time.Time
}

// draft for keyValue
type keyValue struct {
	f imgMeta
	t rectCoord
	d objectDuration
}
