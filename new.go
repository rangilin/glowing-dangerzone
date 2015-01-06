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

	posts := bc.dir + "/_posts"
	if err := os.Mkdir(posts, os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", posts)
	}
	layouts := bc.dir + "/_layouts"
	if err := os.Mkdir(layouts, os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", layouts)
	}
	return nil
}
