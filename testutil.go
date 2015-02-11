package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// assertFileContains assert whether specified file contains specified string
func assertFileContains(t *testing.T, path string, substr string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("Unable to read file %s", path)
	}

	if !strings.Contains(string(content), substr) {
		t.Fatalf("Expect file content contains \n%s\n, but not, content is \n%s\n",
			substr,
			content)
	}
}

// create a temporary folder for testing
func createTmpFolder(t *testing.T) string {
	tmp := filepath.Join(os.TempDir(), "glowing_dangerzone_tmp")
	if _, err := os.Stat(tmp); os.IsNotExist(err) {
		os.Mkdir(tmp, os.ModePerm)
	}

	dir, err := ioutil.TempDir(tmp, "test_")
	if err != nil {
		t.Fatalf("Unable to create temp dir: %s, error: %s", dir, err)
	}
	return dir
}

func testDataPath(subpaths ...string) string {
	relative := filepath.Join(append([]string{"testdata"}, subpaths...)...)
	absolute, _ := filepath.Abs(relative)
	return absolute
}
