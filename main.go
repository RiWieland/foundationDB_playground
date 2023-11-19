package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	/*
	  "github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	  "github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	*/)

// To do:
// - query function: add methods: all and index
// - writing: int to string (for example )
// to check: initImgObj: build function that change object?
var imgSub subspace.Subspace
var rectSub subspace.Subspace

//var processedSub subspace.Subspace

func main() {

	var f img

	//file_path := "2023-10-12T16:02:32.342Z_18:34:02.123Z_cam1.mp4"
	file_path := "2023-10-12T18:34:00.000Z_18:34:02.123Z_cam1.mp4"
	fn := filepath.Base(file_path)

	imgMeta := f.initImgObj(fn)
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

	//  frame manipulation on main:

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
	fdbInst.clearSub(fileDir)

	// Prefex in Subs: first Element of Tuple of the key
	imgSub = fileDir.Sub("img")
	rectSub = fileDir.Sub("rect")

	// Data Model:
	// - key for the rectangle will be the coordinates
	// - no value needed

	seconds := time.Duration(10) * time.Second

	key, _ := fdbInst.writeRect(coor)
	keyN, _ := fdbInst.writeRect(coorN)

	keyImg, _ := fdbInst.writeImgWithCoor(imgMeta, seconds, coor)
	fmt.Println("return coor-key", key)
	fmt.Println("return coor-key", keyN)

	fmt.Println("return fdb-key", keyImg)

	t, _ := fdbInst.queryRectSub()
	fmt.Println(t)
	i, _ := fdbInst.queryImgSub()
	fmt.Println(i)

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
	idx int64
	x0  int64
	y0  int64
	x1  int64
	y1  int64
}

// duration when the object is visible
type objectDuration struct {
	start time.Time
	end   time.Time
}

// draft for keyValue
type keyValue struct {
	f img
	t rectCoord
	d objectDuration
}
