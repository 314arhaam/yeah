package iotools

import (
	// "fmt"
	// "sync"
	"os"
	// "flag"
	// "strconv"
	// "strings"
)

func MakeFixSizeFile(fname string, size int) {
	sample := []byte("Hello, World!\n")
	sampleSize := len(sample)
	file, err := os.Create(fname)
	if err != nil {
		return
	}
	defer file.Close()
	currentSize := 0
	for {
		currentSize += sampleSize
		if currentSize > size {
			sample = sample[:sampleSize-(currentSize-size)-1]
			currentSize = size
		}
		if _, err := file.Write(sample); err != nil {
			return
		}
		if currentSize == size {
			if _, err := file.Write([]byte("\n")); err != nil {
				return
			}
			return
		}
	}
}