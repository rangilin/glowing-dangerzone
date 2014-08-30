package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// Test "glowing-dangerzone create" will create directory and files for blog
func TestCreate(t *testing.T) {
	path := fakeExecutionPath(t)
	os.Args = []string{path, "create"}
	main()

	dir := filepath.Dir(path)
	assertDirExist(t, dir+"/_markdowns")
	assertDirExist(t, dir+"/_engine")
}

// make a fake execution path which is under temporary folder
func fakeExecutionPath(t *testing.T) string {
	path, err := ioutil.TempDir("", "glowing_dangerzone_test_")
	if err != nil {
		t.Fatalf("Unable to create temp dir: %s, error: %s", path, err)
	}
	return filepath.Join(path, filepath.Base(os.Args[0]))
}

func assertDirExist(t *testing.T, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("%s should exist, but not.", path)
	}

}
