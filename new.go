package main

import (
	"fmt"
	"os"
)

// Return a BlogCreator that will create blog engine files in the directory
func NewBlogCreator(dir string) BlogCreator {
	return BlogCreator{dir}
}

// A BlogCreator create necessary files and folder structure for later use.
type BlogCreator struct {
	// directory to put blog engine files
	dir string
}

func (bc BlogCreator) Create() error {
	if isEmpty, _ := isDirEmpty(bc.dir); !isEmpty {
		return fmt.Errorf("%s is not empty", bc.dir)
	}

	showMessage("create _markdowns folder...")
	if err := os.Mkdir(bc.dir+"/_markdowns", os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", bc.dir+"/_markdowns")
	}
	showMessage("create _engine folder...")
	if err := os.Mkdir(bc.dir+"/_engine", os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", bc.dir+"/_engine")
	}
	showMessage("Done ! :)")
	return nil
}
