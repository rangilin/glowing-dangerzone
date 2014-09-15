package main

// -------------------------------------------------- tests for 'new' command
import (
	"os"
	"strings"
	"testing"
)

// New command should create layout in current directory
func TestNew(t *testing.T) {
	dir := createTmpFolder(t)

	createBlogLayout(dir)

	assertFilePathExist(t, dir+"/_markdowns")
	assertFilePathExist(t, dir+"/_engine")
}

func TestNewWhenCurrentFolderIsNotEmpty(t *testing.T) {
	dir := createTmpFolder(t)
	os.Mkdir(dir+"/_whatever", os.ModePerm)

	err := createBlogLayout(dir)

	assertErrorDueToNonEmptyDir(t, err)
}

func assertErrorDueToNonEmptyDir(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("There should be an error, but no error")
	}
	if !strings.HasSuffix(err.Error(), "is not empty") {
		t.Fatalf("Error should due to non empty dir, but due to %s", err.Error())
	}
}
