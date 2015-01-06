package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func isDirEmpty(root string) (bool, error) {
	info, err := os.Lstat(root)
	if err != nil {
		return false, fmt.Errorf("Unable to get file info of %s", root)
	}
	if !info.IsDir() {
		return false, fmt.Errorf("%s is not a directory", root)
	}

	entries := []string{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if path != root {
			entries = append(entries, path)
		}
		return nil
	}
	filepath.Walk(root, walkFn)
	return len(entries) == 0, nil
}
