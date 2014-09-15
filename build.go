package main

import (
	"fmt"
	"os"
)

func buildBlog(dir string) error {
	buildDir := dir + "/blog"

	showMessage("Preparing...")
	os.RemoveAll(buildDir)
	os.Mkdir(buildDir, os.ModePerm)

	file, err := os.Create(buildDir + "/index.html")
	if err != nil {
		return fmt.Errorf("Unable to build blog due to: %s", err.Error())
	}
	file.WriteString("<html><body><h1>Blog</h1></body></html>")
	showMessage("Done :)")
	return nil
}
