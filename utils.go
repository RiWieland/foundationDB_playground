package main

import (
	"fmt"
	"path/filepath"
	"time"
)

func (f file) extractFileMeta(path string) file {
	fn := filepath.Base(path)
	f.object = filepath.Ext(fn)
	//f.startTime = fn[:14]
	//f.endTime = fn[:14]
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
	return f
}
