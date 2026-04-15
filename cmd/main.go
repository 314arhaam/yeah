package main

import (
	"fmt"
	"sync"
	"local/yeah/internal/iotools"
	"time"
)

func main() {
	defer func(t time.Time) {
		fmt.Println("Elapsed Time (ms):", time.Since(t).Milliseconds())
	}(time.Now())
	var cli iotools.CLIArgs
	err := cli.Parse()
	if err != nil {
		fmt.Println("Main: Error @ Parse Args:", err.Error())
		return
	}
	if !(*cli.Parallel) {
		var w sync.WaitGroup
		errChan := make(chan string, len(cli.FileData))
		for _, val := range cli.FileData {
			w.Add(1)
			go func(fname string, fsize int, e chan <- string, wg *sync.WaitGroup) {
				defer wg.Done()
				if err := iotools.MakeFixSizeFile(fname, fsize); err != nil {
					e <- err.Error()
				}
				e <- ""
				return
			}(val.FileName, val.FileSize, errChan, &w)
		}
		w.Wait()
		close(errChan)
		for msg := range errChan {
			if msg != "" {
				fmt.Println("Main: Error @ goroutine `MakeFixedSizeFile()`:", msg)
			}
		}
	} else {
		for _, val := range cli.FileData {
			if err := iotools.MakeFixSizeFile(val.FileName, val.FileSize); err != nil {
				fmt.Println("Main: Error @ Synchron `MakeFixedSizeFile()`:", err.Error())
			}
		}
	}
}
