package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

func TestBuildGeneratingNecessaryFiles(t *testing.T) {
	testDataDir := testDataPath("build", "test_generate_files")
	output := createTmpFolder(t)

	NewBlogBuilder(testDataDir).Build(output)

	postDir := filepath.Join(output, "test-post")
	assertFilePathExist(t, postDir)

	postIndex := filepath.Join(postDir, "index.html")
	assertFilePathExist(t, postIndex)
}

func TestBuildGeneratePostFiles(t *testing.T) {
	testDataDir := testDataPath("build", "test_generate_posts")
	output := createTmpFolder(t)

	NewBlogBuilder(testDataDir).Build(output)

	content, _ := ioutil.ReadFile(filepath.Join(output, "test-post", "index.html"))

	if !strings.Contains(string(content), `<meta http-equiv="X-UA-Compatible" content="IE=edge">`) {
		t.Fatalf("No base template in post file")
	}

	if !strings.Contains(string(content), "This is test post content") {
		t.Fatalf("No post in post file")
	}
}

func TestBuildBlogIndexPage(t *testing.T) {
	testDataDir := testDataPath("build", "test_generate_index")
	output := createTmpFolder(t)

	NewBlogBuilder(testDataDir).Build(output)

	bytes, _ := ioutil.ReadFile(filepath.Join(output, "index.html"))
	content := string(bytes)

	if !strings.Contains(content, `<meta http-equiv="X-UA-Compatible" content="IE=edge">`) {
		t.Fatalf("No base template in blog index file")
	}
	if !strings.Contains(content, "Test Post") {
		t.Fatalf("No post in blog index file")
	}
}
