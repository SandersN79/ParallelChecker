package main

import (
	"ParallelChecker/core"
	"fmt"
	"os"
)

func errorExit(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func main() {
	inputs, err := getCliInputs()
	if err != nil {
		errorExit(err)
	}
	err = core.DuplicateCheck(inputs.filePaths) // Check csv files for duplicate codes by inputting .csv file names
	if err != nil {
		errorExit(err)
	}
	fmt.Println("No Duplicates Found...")
}
