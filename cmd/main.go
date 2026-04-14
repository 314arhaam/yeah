package main

import (
	"fmt"
	"sync"
	"local/yeah/internal/iotools"
	"time"
)

func main() {
	startTime := time.Now()
	var cli iotools.CLIArgs
	err := cli.Parse()
	if err != nil {
		fmt.Println("Main: Error", err.Error())
		return
	}
	if !(*cli.Parallel) {
		var w sync.WaitGroup
		for _, val := range cli.FileData {
			w.Add(1)
			go func(fname string, fsize int, wg *sync.WaitGroup) {
				defer wg.Done()
				iotools.MakeFixSizeFile(fname, fsize)
				return
			}(val.FileName, val.FileSize, &w)
		}
		w.Wait()
	} else {
		for _, val := range cli.FileData {
			iotools.MakeFixSizeFile(val.FileName, val.FileSize)
		}
	}
	fmt.Println("Elapsed Time: ", time.Since(startTime).Milliseconds())
}
