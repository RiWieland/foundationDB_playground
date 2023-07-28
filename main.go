package main

import (
	"fmt"
	"log"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
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

	var f file

	//file_path := "2023-10-12T16:02:32.342Z_18:34:02.123Z_cam1.mp4"
	file_path := "2023-10-12T18:34:00.000Z_18:34:02.123Z_cam1.mp4"
	fMeta := f.extractFileMeta(file_path)
	fmt.Println(fMeta.endTime.String())

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
	db := initFdb()
	//db.instance = initFdb()
	//db.instance.Options().SetTransactionTimeout(60000) // 60,000 ms = 1 minute
	//db.instance.Options().SetTransactionRetryLimit(100)

	// add meta to file subDirectory:
	fileDir, err := directory.CreateOrOpen(db, []string{"fileDir"}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileDir)
	//rawSub := fileDir.Sub("rawVideo")
	//proSub := fileDir.Sub("processedVideo")
	//db.addDirectorySub("rawVideo")
	//fmt.Println(rawSub)

}

// write for tag in video
type file struct {
	startTime time.Time
	endTime   time.Time
	path      string
	object    string
	fileType  string
}

// coordinates where object is marked
type objectCoord struct {
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
	t objectCoord
	d objectDuration
}
