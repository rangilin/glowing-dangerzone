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

	baseTemplate, err := os.Create(filepath.Join(layouts, "base.tmpl"))
	if err != nil {
		return fmt.Errorf("Unable to create base.tmpl due to %v", err)
	}
	baseTemplate.WriteString(BaseTemplateContent)

	postTemplate, err := os.Create(filepath.Join(layouts, "post.tmpl"))
	if err != nil {
		return fmt.Errorf("Unable to create post.tmpl due to %v", err)
	}
	postTemplate.WriteString(PostTemplateContent)

	indexTemplate, err := os.Create(filepath.Join(layouts, "index.tmpl"))
	if err != nil {
		return fmt.Errorf("Unable to create index.tmpl due to %v", err)
	}
	indexTemplate.WriteString(IndexTemplateContent)
	return nil
}
