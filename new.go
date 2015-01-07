package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	if isEmpty, _ := IsDirEmpty(bc.dir); !isEmpty {
		return fmt.Errorf("%s is not empty", bc.dir)
	}

	posts := filepath.Join(bc.dir, PostsDirName)
	if err := os.Mkdir(posts, os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", posts)
	}
	layouts := filepath.Join(bc.dir, LayoutsDirName)
	if err := os.Mkdir(layouts, os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", layouts)
	}
	return nil
}
