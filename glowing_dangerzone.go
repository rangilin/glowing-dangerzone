package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	dir := getCurrentDir()
	cmd := getSubCommand()
	switch cmd {
	case "new":
		NewBlogCreator(dir).Create()
	case "build":
		NewBlogBuilder(dir + "/blog").Build()
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

func getCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
