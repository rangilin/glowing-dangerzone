package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	var dir string
	var err error
	dir, err = getCurrentDir()
	if err != nil {
		logErrorAndShutdown(err)
	}

	cmd := os.Args[1]
	if cmd == "new" {
		err = createBlogLayout(dir)
	} else if cmd == "build" {
		err = BlogBuilder{dir + "/blog"}.Build()
	}

	if err != nil {
		logErrorAndShutdown(err)
	}
}

func getCurrentDir() (dir string, err error) {
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	return
}

func logErrorAndShutdown(err error) {
	log.Fatal(err)
	os.Exit(1)
}
