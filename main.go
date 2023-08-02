package main

import (
	"fmt"
	"log"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
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
// - transfer data between routines -> channnel /chan
// - make sure every person only got one account
// - implement restrictions:
// -  - no negative vales
// - overwriting keys in key-value store OR timestamp?

var imgSub subspace.Subspace
var rectSub subspace.Subspace

//var processedSub subspace.Subspace

func main() {

	var f file

	//file_path := "2023-10-12T16:02:32.342Z_18:34:02.123Z_cam1.mp4"
	file_path := "2023-10-12T18:34:00.000Z_18:34:02.123Z_cam1.mp4"
	fMeta := f.extractImgMeta(file_path)
	fmt.Println(fMeta.time.String())

	/* Put this in place for frame manipulation:
	// Image Manipulation, External Model:
	EditImg := readImg("test.jpg")

	coor := objectCoord{
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

	// clear:
	_, err = fdbInst.instance.Transact(func(tr fdb.Transaction) (interface{}, error) {
		tr.ClearRange(fileDir)
		return nil, nil
	})

	// Data Model:
	// - key for the rectangle will be the coordinates
	// - no value needed
	testCoor := rectCoord{1, 3, 4, 5}
	seconds := time.Duration(10) * time.Second

	fdbInst.writeRect(testCoor)
	fdbInst.writeFileCoor(fMeta, testCoor, seconds)
}

// write for tag in video
type file struct {
	path     string
	fileType string
	time     time.Time
	rect     string
}

// coordinates where object is marked
type rectCoord struct {
	x0 int
	y0 int
	x1 int
	y1 int
}

// duration when the object is visible
type objectDuration struct {
	start time.Time
	end   time.Time
}

// draft for keyValue
type keyValue struct {
	f file
	t rectCoord
	d objectDuration
}
