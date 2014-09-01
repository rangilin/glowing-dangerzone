package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createBlogLayout() error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return fmt.Errorf("Unable to get file current file path")
	}
	if err := os.Mkdir(dir+"/_markdowns", os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", dir+"/_markdowns")
	}
	if err := os.Mkdir(dir+"/_engine", os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", dir+"/_engine")
	}
	return nil
}
