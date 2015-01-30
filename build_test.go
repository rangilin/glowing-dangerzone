package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildWillCreateBlogFolder(t *testing.T) {
	dir := createTmpFolder(t)
	output := filepath.Join(createTmpFolder(t), BuildDirName)

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

func TestBuildGeneratePostFiles(t *testing.T) {
	testDataDir := testDataPath("build", "test_generate_posts")
	output := filepath.Join(createTmpFolder(t), "blog")

	NewBlogBuilder(testDataDir).Build(output)

	postDir := filepath.Join(output, "test-build-1")
	index := filepath.Join(postDir, "index.html")
	assertFilePathExist(t, postDir)
	assertFilePathExist(t, index)

	// post content
	assertFileContains(t, index, "This is test build 1 content")
	// base template
	assertFileContains(t, index, `<meta http-equiv="X-UA-Compatible" content="IE=edge">`)
}

func TestBuildGenerateIndexPage(t *testing.T) {

}
