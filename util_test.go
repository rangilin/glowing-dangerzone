package main

import (
	"os"
	"testing"
)

func TestIsDirEmpty(t *testing.T) {
	emptyDir := createTmpFolder(t)
	if isEmpty, _ := isDirEmpty(emptyDir); !isEmpty {
		t.Fatalf("Dir %s is empty, but return false", emptyDir)
	}

	nonEmptyDir := createTmpFolder(t)
	os.Mkdir(nonEmptyDir+"/_whatever", os.ModePerm)
	if isEmpty, _ := isDirEmpty(nonEmptyDir); isEmpty {
		t.Fatalf("Dir %s is not empty, but return true", nonEmptyDir)
	}
}
