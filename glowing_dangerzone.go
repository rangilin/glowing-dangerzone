package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	dir := getCurrentDir()
	cmd := getSubCommand()
	if cmd == "new" {
		createBlogLayout(dir)
	} else if cmd == "build" {
		BlogBuilder{dir + "/blog"}.Build()
	} else {
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

func getCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
