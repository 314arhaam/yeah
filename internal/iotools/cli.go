package iotools

import (
	"fmt"
	"flag"
	"strconv"
	"strings"
)

type FileData struct {
	FileName	string
	FileSize	int
}

type CLIArgs struct {
	fileSize	*string
	// exported
	Parallel	*bool
	FileData	[]FileData
}

func getFileData(s string) ([]FileData, error) {
	argsSlice := strings.Split(s, ",")
	fds := make([]FileData, len(argsSlice))
	for i, val := range argsSlice {
		trimVal := strings.Trim(val, " ")
		// 
		switch strings.ToUpper(string(trimVal[len(trimVal)-1])) {
		case "K":
			tmp, err := strconv.Atoi(trimVal[:len(trimVal)-1])
			if err != nil {
				return nil, fmt.Errorf("Error in `AtoiSlice`: Value %s cannot be parsed as int.", val)
			}
			tmp = tmp*1024
			fds[i].FileSize = tmp
		case "M":
			tmp, err := strconv.Atoi(trimVal[:len(trimVal)-1])
			if err != nil {
				return nil, fmt.Errorf("Error in `AtoiSlice`: Value %s cannot be parsed as int.", val)
			}
			tmp = tmp*1024*1024
			fds[i].FileSize = tmp
		case "G":
			return nil, fmt.Errorf("Error in `AtoiSlice`: GB not implemented.")
		default:
			tmp, err := strconv.Atoi(trimVal)
			if err != nil {
				return nil, fmt.Errorf("Error in `AtoiSlice`: Value %s cannot be parsed as int.", val)
			}
			fds[i].FileSize = tmp
		}
		fname := ("yeah_" + strconv.Itoa(i) + "-" + val + ".txt")
		fds[i].FileName = fname
	}
	return fds, nil
}

func (cli *CLIArgs) Parse() error {
	cli.Parallel = flag.Bool("s", false, "Synchron mode")
	cli.fileSize = flag.String("f", "100", "Size of file(s) in bytes, comma separated")
	flag.Parse()
	res, err := getFileData(*cli.fileSize)
	if err != nil {
		return fmt.Errorf("Error in `(cli *CLIArgs) Parse()`:", err)
	}
	cli.FileData = res
	return nil
}