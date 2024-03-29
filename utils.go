package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// Function takes the path and initialize imgObject out of it
func (f img) initImgObj(fn string) img {

	f.path = fn
	f.fileType = filepath.Ext(fn)
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

func getFramesPerSec(startSec int, endSec int) [2]int {
	var FrameArray [2]int
	FrameArray[0] = int(float64(startSec) * 25.1)
	FrameArray[1] = int((float64(endSec) * 25.1))
	// 30, 60 oder gar 120
	return FrameArray
}

func exampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {

	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

func extractFrames(input_path string, output_path string, start_sec int, end_sec int) {
	target_frames := getFramesPerSec(start_sec, end_sec)

	for i := target_frames[0]; i < target_frames[1]; i++ {

		reader := exampleReadFrameAsJpeg(input_path, (int(i)))
		img, err := imaging.Decode(reader)
		if err != nil {
			fmt.Println("ERROR")
		}

		str := strconv.Itoa(i)
		target_path := output_path + "out" + str + ".jpeg"
		err = imaging.Save(img, target_path)
		if err != nil {
			fmt.Println("ERROR")
		}
	}
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func convertToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(s, "is not an inteer.")
	}
	return n
}
