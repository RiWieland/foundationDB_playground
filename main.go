package main

import (
	"fmt"

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
