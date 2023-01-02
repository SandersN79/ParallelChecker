package utils

import (
	"os"
	"path/filepath"
)

// walkDirectory recursively finds all csv files in an input directory
func walkDirectory(rootPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".csv" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// listFiles lists csv files located in a directory
func listFiles(path string) ([]string, error) {
	var files []string
	dir, err := os.Open(path)
	if err != nil {
		return files, err
	}
	defer dir.Close()
	dirFiles, err := dir.Readdir(-1)
	if err != nil {
		return files, err
	}
	for _, f := range dirFiles {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".csv" {
			files = append(files, path+f.Name())
		}
	}
	return files, nil
}

// GetFiles in a path either recursively or non-recursively
func GetFiles(path string, recur bool) ([]string, error) {
	if recur {
		return walkDirectory(path)
	}
	return listFiles(path)
}
