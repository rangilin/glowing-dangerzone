package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
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
		NewBlogBuilder(dir).Build(filepath.Join(dir, BuildDirName))
	case "post":
		title := parseCreatePostTitle()
		NewPostCreator(filepath.Join(dir, PostsDirName)).Create(title)
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

func parseCreatePostTitle() string {
	var title = ""
	flagSet := flag.NewFlagSet("post", flag.ExitOnError)
	flagSet.StringVar(&title, "title", "", "post title")
	flagSet.Parse(os.Args[2:])

	title = strings.TrimSpace(title)
	if title == "" {
		log.Fatalf("post title is required")
	}
	return title
}
