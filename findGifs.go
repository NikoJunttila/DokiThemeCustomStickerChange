package main

import (
	"os"
	"path/filepath"
)

func getAllFileNames(folderPath string) ([]string, error) {
	var fileNames []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileNames = append(fileNames, info.Name())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileNames, nil
}
