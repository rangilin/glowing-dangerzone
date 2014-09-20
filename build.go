package main

import (
	"fmt"
	"os"
)

type BlogBuilder struct {
	dir string // directory to put generated files
}

func (b BlogBuilder) Build() error {
	showMessage("Preparing...")

	b.cleanup()

	file, err := os.Create(b.dir + "/index.html")
	if err != nil {
		return fmt.Errorf("Unable to build blog due to: %s", err.Error())
	}
	file.WriteString("<html><body><h1>Blog</h1></body></html>")
	showMessage("Done :)")
	return nil
}

func (b BlogBuilder) cleanup() {
	os.RemoveAll(b.dir)
	os.Mkdir(b.dir, os.ModePerm)
}
