package main

import (
	"os"
	"testing"
)

func TestBuildWillCreateBlogFolder(t *testing.T) {
	dir := createTmpFolder(t)

	buildBlog(dir)

	assertFilePathExist(t, dir+"/blog")
}

func TestBlogFolderWillBeDeletedBeforeBuild(t *testing.T) {
	dir := createTmpFolder(t)
	os.Mkdir(dir+"/blog", os.ModePerm)
	os.Create(dir + "/blog/delete_me")

	buildBlog(dir)

	if _, err := os.Stat(dir + "/blog/delete_me"); !os.IsNotExist(err) {
		t.Fatalf("Should delete exist blog folder before build")
	}
}

func TestBuildBlogWillGenerateIndexPage(t *testing.T) {
	dir := createTmpFolder(t)

	buildBlog(dir)

	assertFilePathExist(t, dir+"/blog/index.html")
}
