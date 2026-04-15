package iotools

import (
	"os"
	"fmt"
)

func MakeFixSizeFile(fname string, size int) error {
	sample := []byte("Hello, World!\n")
	sampleSize := len(sample)
	file, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("Error in `MakeFixSizeFile`: Cannot create file: os.Create")
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
			return fmt.Errorf("Error in `MakeFixSizeFile`: Cannot write data to file.")
		}
		if currentSize == size {
			if _, err := file.Write([]byte("\n")); err != nil {
				return fmt.Errorf("Error in `MakeFixSizeFile`: Cannot write data to file. Last chunk.")
			}
			return nil
		}
	}
}