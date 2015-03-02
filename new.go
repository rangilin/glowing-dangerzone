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

	assets := filepath.Join(bc.dir, AssetsDirName)
	if err := os.Mkdir(assets, os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", assets)
	}

	templates := [...][2]string{
		[2]string{"base.tmpl", BaseTemplateContent},
		[2]string{"post.tmpl", PostTemplateContent},
		[2]string{"index.tmpl", IndexTemplateContent},
		[2]string{"feeds.xml", FeedsXMLContent},
	}

	for _, template := range templates {
		name := template[0]
		content := template[1]
		t, err := os.Create(filepath.Join(layouts, name))
		if err != nil {
			return fmt.Errorf("Unable to create %s due to %v", name, err)
		}
		_, err = t.WriteString(content)
		if err != nil {
			return err
		}
	}
	return nil
}
