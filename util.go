package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
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

	s := strings.ToLower(str)
	s = strings.Replace(s, "'", "", -1)
	s = reg.ReplaceAllString(s, " ")

	chars := make([]string, 0)
	for _, v := range strings.Split(s, " ") {
		if len(v) != 0 {
			chars = append(chars, v)
		}
	}
	return strings.Join(chars, "-")
}

func CopyFile(source string, destination string) error {
	content, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(destination, content, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// UTF8Slice just like ordinary sub string function, but count string by
// rune instead of byte.
func UTF8Slice(str string, startIdx int, length int) (string, error) {

	if startIdx < 0 {
		return "", fmt.Errorf("Start index out of bound")
	}
	if length < 0 {
		return "", fmt.Errorf("Length invalid")
	}

	idx := 0
	slicelen := 0
	var b bytes.Buffer
	for len(str) > 0 && length > 0 {
		r, size := utf8.DecodeRuneInString(str)
		if idx >= startIdx {
			slicelen++
			b.WriteRune(r)
		}

		str = str[size:]
		if idx < startIdx && len(str) == 0 {
			return "", fmt.Errorf("Start index out of bound")
		}

		if slicelen >= length || len(str) == 0 {
			return b.String(), nil
		}
		idx++
	}
	return "", nil
}
