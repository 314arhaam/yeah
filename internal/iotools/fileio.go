package iotools

import (
	"os"
	"fmt"
	"bytes"
)

func MakeFixSizeFileLinear(fname string, size int) error {
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

func MakeFixSizeFile(fname string, size int) error {
	sample := []byte("Hello, World!\n")
	file, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("Error in `MakeFixSizeFile`: Cannot create file: os.Create")
	}
	defer file.Close()
	currentSize := 0
	for {
		sampleSize := len(sample)
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
		sample = bytes.Repeat(sample, 2)
	}
}