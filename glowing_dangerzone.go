package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	cmd := os.Args[1]
	var err error
	if cmd == "new" {
		err = createBlogLayout()
	} else if cmd == "build" {
		err = buildBlog()
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
