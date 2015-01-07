package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Date Format for ISO 8601
const ISO8601Date = "2006-01-02"

func IsDirEmpty(root string) (bool, error) {
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

// Prettify string into url/filepath friendly string
func Prettify(str string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")

	s := strings.Replace(str, "'", "", -1)
	s = reg.ReplaceAllString(s, " ")

	chars := make([]string, 0)
	for _, v := range strings.Split(s, " ") {
		if len(v) != 0 {
			chars = append(chars, v)
		}
	}
	return strings.Join(chars, "-")
}
