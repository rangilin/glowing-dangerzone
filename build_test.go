package main

import (
	"os"
	"testing"
)

func TestBuildWillCreateBlogFolder(t *testing.T) {
	dir := createTmpFolder(t)

	NewBlogBuilder(dir + "/blog").Build()

	assertFilePathExist(t, dir+"/blog")
}

func TestBlogFolderWillBeDeletedBeforeBuild(t *testing.T) {
	dir := createTmpFolder(t)
	os.Mkdir(dir+"/blog", os.ModePerm)
	os.Create(dir + "/blog/delete_me")

	NewBlogBuilder(dir + "/blog").Build()

	if _, err := os.Stat(dir + "/blog/delete_me"); !os.IsNotExist(err) {
		t.Fatalf("Should delete exist blog folder before build")
	}
}

func TestBuildBlogWillGenerateIndexPage(t *testing.T) {
	dir := createTmpFolder(t)

	NewBlogBuilder(dir + "/blog").Build()

	assertFilePathExist(t, dir+"/blog/index.html")
}
