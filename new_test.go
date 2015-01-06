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

	NewBlogCreator(dir).Create()

	assertFilePathExist(t, dir+"/_posts")
	assertFilePathExist(t, dir+"/_layouts")
}

func TestNewWhenCurrentFolderIsNotEmpty(t *testing.T) {
	dir := createTmpFolder(t)
	os.Mkdir(dir+"/_whatever", os.ModePerm)

	err := NewBlogCreator(dir).Create()

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
