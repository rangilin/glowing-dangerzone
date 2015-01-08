package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildWillCreateBlogFolder(t *testing.T) {
	dir := createTmpFolder(t)
	output := filepath.Join(dir, "blog")

	NewBlogBuilder(dir).Build(output)

	assertFilePathExist(t, output)
}

func TestCleanUpBeforeBuild(t *testing.T) {
	dir := createTmpFolder(t)
	output := filepath.Join(dir, "blog")
	deleteme := filepath.Join(output, "delete_me")
	os.Mkdir(output, os.ModePerm)
	os.Create(deleteme)

	NewBlogBuilder(dir).Build(output)

	if _, err := os.Stat(deleteme); !os.IsNotExist(err) {
		t.Fatalf("Should delete exist blog folder before build")
	}
}

func TestBuild(t *testing.T) {

}
