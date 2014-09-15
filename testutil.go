package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// assertFilePathExist assert specified path entry is exist.
func assertFilePathExist(t *testing.T, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("%s should exist, but not.", path)
	}
}

func createTmpFolder(t *testing.T) string {
	dir, err := ioutil.TempDir("", "glowing_dangerzone_test_")
	if err != nil {
		t.Fatalf("Unable to create temp dir: %s, error: %s", dir, err)
	}
	return dir
}
