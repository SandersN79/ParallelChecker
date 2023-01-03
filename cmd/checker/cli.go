package main

import (
	"flag"
	"github.com/SandersN79/parallelChecker/utils"
	"os"
	"strings"
)

// cliInputs stores the read input data from the CLI
type cliInputs struct {
	filePaths []string
}

// getCliInputs parses CLI inputs
func getCliInputs() (*cliInputs, error) {
	var err error
	var folderPath, fileNames string
	var recur bool
	flag.StringVar(&folderPath, "path", "", "folder path containing csv files")
	flag.StringVar(&fileNames, "files", "*", "csv file names in folder path to check")
	flag.BoolVar(&recur, "recursive", false, "if true, recursively walk folder path for csv files")
	flag.Parse()
	if folderPath == "" { // set default testData folderPath if none entered
		folderPath, err = os.Getwd()
		if err != nil {
			return &cliInputs{}, err
		}
		folderPath += "/testData/"
	}
	if folderPath[len(folderPath)-1:] != "/" { // ensure folderPath ends with '/'
		folderPath += "/"
	}
	files := strings.Split(fileNames, ",")
	var filePaths []string
	if len(files) == 0 || fileNames == "*" { // if no files specified; get all files from folderPath
		filePaths, err = utils.GetFiles(folderPath, recur)
		if err != nil {
			return &cliInputs{}, err
		}
	} else {
		for _, f := range files { // else get builder filePaths for each specified file
			filePaths = append(filePaths, folderPath+f)
		}
	}
	return &cliInputs{
		filePaths,
	}, nil
}
