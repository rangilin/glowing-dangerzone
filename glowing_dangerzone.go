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

	cmd := os.Args[1]
	if cmd == "create" {
		createBlogLayout()
	}
}

func createBlogLayout() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	mkdir(dir + "/_markdowns")
	mkdir(dir + "/_engine")
}

func mkdir(dir string) {
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
