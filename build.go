package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func NewBlogBuilder(dir string) BlogBuilder {
	return BlogBuilder{dir, NewPostParser()}
}

// A BlogBuilder that generate static files from posts/layouts
type BlogBuilder struct {
	// where posts, layouts directory exist
	dir string

	// Parser that parse Post from post files
	postParser PostParser
}

// Generate static files to specified directory
func (b BlogBuilder) Build(output string) error {
	os.RemoveAll(output)
	os.Mkdir(output, os.ModePerm)

	//base, _ := template.ParseFiles(filepath.Join(b.dir, LayoutsDirName, "base.tmpl"))

	for _, path := range b.getPostPaths() {
		post := b.postParser.Parse(path)

		postDir := filepath.Join(output, Prettify(post.Title()))
		err := os.Mkdir(postDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Unable to create post folder %s", postDir)
		}

		index := filepath.Join(postDir, "index.html")
		file, err := os.Create(index)
		if err != nil {
			return fmt.Errorf("Unable to create file %s", index)
		}
		file.WriteString(post.HtmlContent())
	}
	return nil
}

func (b BlogBuilder) getPostPaths() []string {
	postDir := filepath.Join(b.dir, PostsDirName)
	paths, _ := filepath.Glob(postDir + string(os.PathSeparator) + "*")
	return paths
}
