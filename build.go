package main

import (
	"os"
	"path/filepath"
)

func NewBlogBuilder(dir string) BlogBuilder {
	return BlogBuilder{dir}
}

// A BlogBuilder that generate static files from posts/layouts
type BlogBuilder struct {
	// where posts, layouts directory exist
	dir string

	//postParser PostParser
}

// Generate static files to specified directory
func (b BlogBuilder) Build(outputPath string) error {
	b.cleanup()

	return nil
}

func (b BlogBuilder) cleanup() {
	output := filepath.Join(b.dir, BuildDirName)
	os.RemoveAll(output)
	os.Mkdir(output, os.ModePerm)
}
