package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (f file) extractImgMeta(path string) file {
	fn := filepath.Base(path)
	f.fileType = filepath.Ext(fn)
	f.path = fn
	//f.startTime = fn[:14]
	//f.endTime = fn[:14]
	/*
		startTimeString := fn[:24]
		endtimeString := fn[:11] + fn[26:38]

		startTime, err := time.Parse(time.RFC3339, startTimeString)
		if err != nil {
			fmt.Println("Error while parsing date :", err)
		}
		endTime, err := time.Parse(time.RFC3339, endtimeString)
		if err != nil {
			fmt.Println("Error while parsing date :", err)
		}

		f.startTime = startTime
		f.endTime = endTime
	*/
	return f
}

func createOutputFile(path string) *os.File {
	// Add Point Slice
	out, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return out
}
