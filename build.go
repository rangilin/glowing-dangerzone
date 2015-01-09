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
func (b BlogBuilder) Build(output string) error {
	os.RemoveAll(output)
	os.Mkdir(output, os.ModePerm)

	for _, path := range b.getPostPaths() {
		post := b.postParser.Parse(path)
		os.Mkdir(filepath.Join(output, Prettify(post.Title())), os.ModePerm)
	}
	return nil
}

func (b BlogBuilder) getPostPaths() []string {
	postDir := filepath.Join(b.dir, PostsDirName)
	paths, _ := filepath.Glob(postDir + string(os.PathSeparator) + "*")
	return paths
}
