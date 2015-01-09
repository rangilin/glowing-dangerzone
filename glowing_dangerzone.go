package main

import (
	"log"
	"os"
	"path/filepath"
)

// constant for directory name
const (
	// name of directory contains build result
	BuildDirName = "blog"
	// name of directory contains user post files
	PostsDirName = "posts"
	// name of post file
	PostFileName = "post.md"
	// name of directory contains template files
	LayoutsDirName = "layouts"
)

func main() {

	dir, _ := os.Getwd()
	cmd := getSubCommand()
	switch cmd {
	case "new":
		NewBlogCreator(dir).Create()
	case "build":
		NewBlogBuilder(filepath.Join(dir, BuildDirName)).Build("")
	case "post":
		NewPostCreator(filepath.Join(dir, PostsDirName)).Create("test")
	case "serve":
		RunFileServer(filepath.Join(dir, BuildDirName))
	default:
		log.Fatalf("unknown command %s", cmd)
	}
}

func getSubCommand() string {
	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	return cmd
}
