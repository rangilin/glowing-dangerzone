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
	fmt.Println("[GD] create _markdowns folder...")
	if err := os.Mkdir(dir+"/_markdowns", os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", dir+"/_markdowns")
	}
	fmt.Println("[GD] create _engine folder...")
	if err := os.Mkdir(dir+"/_engine", os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", dir+"/_engine")
	}
	fmt.Println("[GD] Done ! :)")
	return nil
}
