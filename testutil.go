package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// assertFilePathExist assert specified path entry is exist.
func assertFilePathExist(t *testing.T, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("%s should exist, but not.", path)
	}
}

// fakeExecutionPath change current execution path to a temporary directory.
func fakeExecutionPath(t *testing.T) {
	path, err := ioutil.TempDir("", "glowing_dangerzone_test_")
	if err != nil {
		t.Fatalf("Unable to create temp dir: %s, error: %s", path, err)
	}
	os.Args[0] = filepath.Join(path, filepath.Base(os.Args[0]))
}

func createTmpFolder(t *testing.T) string {
	dir, err := ioutil.TempDir("", "glowing_dangerzone_test_")
	if err != nil {
		t.Fatalf("Unable to create temp dir: %s, error: %s", dir, err)
	}
	return dir
}
