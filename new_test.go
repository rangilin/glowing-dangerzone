package main

// -------------------------------------------------- tests for 'new' command
import (
	"testing"
)

// New command should create layout in current directory
func TestNew(t *testing.T) {
	dir := createTmpFolder(t)

	createBlogLayout(dir)

	assertFilePathExist(t, dir+"/_markdowns")
	assertFilePathExist(t, dir+"/_engine")
}
