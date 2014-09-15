package main

import (
	"fmt"
	"os"
)

func createBlogLayout(dir string) error {
	showMessage("create _markdowns folder...")
	if err := os.Mkdir(dir+"/_markdowns", os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", dir+"/_markdowns")
	}
	showMessage("create _engine folder...")
	if err := os.Mkdir(dir+"/_engine", os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", dir+"/_engine")
	}
	showMessage("Done ! :)")
	return nil
}
