package main

import (
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
func (b BlogBuilder) Build(outputPath string) error {
	os.RemoveAll(outputPath)
	os.Mkdir(outputPath, os.ModePerm)
	return nil
}

func (b BlogBuilder) getPostPaths() []string {
	postDir := filepath.Join(b.dir, PostsDirName)
	paths, _ := filepath.Glob(postDir + string(os.PathSeparator) + "*")
	return paths
}
