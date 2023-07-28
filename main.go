package main

import (
	"fmt"
	"log"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
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

var rawSub subspace.Subspace

//var processedSub subspace.Subspace

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

	// add meta to file subDirectory:
	fileDir, err := directory.CreateOrOpen(db, []string{"fileDir"}, nil)
	if err != nil {
		log.Fatal(err)
	}

	rawSub := fileDir.Sub("rawVideo")
	//processedSub := fileDir.Sub("processedVideo")

	// clear:
	_, err = db.Transact(func(tr fdb.Transaction) (interface{}, error) {
		tr.ClearRange(fileDir)
		return nil, nil
	})

	// enter values:
	SCKey := rawSub.Pack(tuple.Tuple{f.path, f.startTime, f.endTime})
	_, err = db.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		tr.Set(SCKey, []byte{})
		return
	})

}

func signup(t fdb.Transactor, studentID, class string) (err error) {
	SCKey := rawSub.Pack(tuple.Tuple{studentID, class})

	_, err = t.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		tr.Set(SCKey, []byte{})
		return
	})
	return
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
