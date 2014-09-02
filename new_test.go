package main

// -------------------------------------------------- tests for 'new' command
import (
	"os"
	"path/filepath"
	"testing"
)

// New command should create layout in current directory
func TestNew(t *testing.T) {
	fakeExecutionPath(t)
	os.Args = append(os.Args, "new")

	main()

	dir := filepath.Dir(os.Args[0])
	assertFilePathExist(t, dir+"/_markdowns")
	assertFilePathExist(t, dir+"/_engine")
}
