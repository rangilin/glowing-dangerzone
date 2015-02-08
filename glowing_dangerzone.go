package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// constants for directory name
const (
	// name of directory contains build result
	BuildDirName = "blog"
	// name of directory contains user post files
	PostsDirName = "posts"
	// name of directory contains user post files
	AssetsDirName = "assets"
)

// constants for file name
const (
	// name of post file
	PostFileName = "post.md"
	// name of directory contains template files
	LayoutsDirName = "layouts"
	// name of base template file
	BaseTemplateName = "base.tmpl"
	// name of blog index template file
	IndexTemplateName = "index.tmpl"
	// name of post template file
	PostTemplateName = "post.tmpl"
)

func main() {
	var err error

	dir, _ := os.Getwd()
	conf := getConfiguration()
	cmd := getSubCommand()
	switch cmd {
	case "new":
		err = NewBlogCreator(dir).Create()
	case "build":
		err = NewBlogBuilder(conf, dir).Build(filepath.Join(dir, BuildDirName))
	case "post":
		title := parseCreatePostTitle()
		err = NewPostCreator(filepath.Join(dir, PostsDirName)).Create(title)
	case "serve":
		RunFileServer(filepath.Join(dir, BuildDirName))
	default:
		log.Fatalf("unknown command %s", cmd)
	}

	if err != nil {
		log.Fatal("Errors : ", err)
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
