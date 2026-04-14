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
		tmp, err := strconv.Atoi(strings.Trim(val, " "))
		if err != nil {
			return nil, fmt.Errorf("Error in `AtoiSlice`: Value %s cannot be parsed as int.", val)
		}
		fname := ("yeah_" + strconv.Itoa(i) + "-" + val + ".txt")
		fds[i].FileSize = tmp
		fds[i].FileName = fname
	}
	return fds, nil
}

func (cli *CLIArgs) Parse() error {
	cli.Parallel = flag.Bool("s", false, "If works in parallel on files")
	cli.fileSize = flag.String("f", "100", "Size of file(s) in bytes, comma separated")
	flag.Parse()
	res, err := getFileData(*cli.fileSize)
	if err != nil {
		return fmt.Errorf("Error in `(cli *CLIArgs) Parse()`:", err)
	}
	cli.FileData = res
	return nil
}