package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildWillCreateBlogFolder(t *testing.T) {
	t.Parallel()

	dir := createTmpFolder(t)
	output := filepath.Join(createTmpFolder(t), BuildDirName)

	NewBlogBuilder(getConfiguration(), dir).Build(output)

	if _, err := os.Stat(output); os.IsNotExist(err) {
		t.Fatalf("Blog folder should be created at %s, but not.", output)
	}
}

func TestCleanUpBeforeBuild(t *testing.T) {
	t.Parallel()

	dir := createTmpFolder(t)
	output := filepath.Join(dir, "blog")

	deleteme := filepath.Join(output, "delete_me")
	os.Mkdir(output, os.ModePerm)
	os.Create(deleteme)

	NewBlogBuilder(getConfiguration(), dir).Build(output)

	if _, err := os.Stat(deleteme); !os.IsNotExist(err) {
		t.Fatalf("Should delete exist blog folder before build")
	}
}

func TestBuildGeneratingNecessaryFiles(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_generate_files")
	output := createTmpFolder(t)

	NewBlogBuilder(getConfiguration(), testDataDir).Build(output)

	postDir := filepath.Join(output, "test-post")
	if _, err := os.Stat(postDir); os.IsNotExist(err) {
		t.Fatalf("Post folder should be created at %s, but not", postDir)
	}

	postIndex := filepath.Join(postDir, "index.html")
	if _, err := os.Stat(postIndex); os.IsNotExist(err) {
		t.Fatalf("Post index file should be created at %s, but not", postIndex)
	}
}

func TestBuildGeneratePostFiles(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_generate_posts")
	output := createTmpFolder(t)

	NewBlogBuilder(getConfiguration(), testDataDir).Build(output)

	bytes, _ := ioutil.ReadFile(filepath.Join(output, "test-post", "index.html"))
	content := string(bytes)

	if !strings.Contains(content, `<meta http-equiv="X-UA-Compatible" content="IE=edge">`) {
		t.Fatalf("No base template in post file")
	}

	if !strings.Contains(content, "<p>This is test post content</p>") {
		t.Fatalf("No post in post file")
	}
}

func TestBuildShouldCopyPostFiles(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_generate_posts")
	output := createTmpFolder(t)

	NewBlogBuilder(getConfiguration(), testDataDir).Build(output)

	path := filepath.Join(output, "test-post", "test.txt")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("%s should be copied to the built result, but not.", path)
	}
}

func TestBuildShouldCopyPostFilesRecursively(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_generate_posts")
	output := createTmpFolder(t)

	NewBlogBuilder(getConfiguration(), testDataDir).Build(output)

	path := filepath.Join(output, "test-post", "test", "test.txt")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("%s should be copied to the built result, but not.", path)
	}
}

func TestBuildShouldNotCopyPostMarkdownFile(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_generate_posts")
	output := createTmpFolder(t)

	NewBlogBuilder(getConfiguration(), testDataDir).Build(output)

	path := filepath.Join(output, "test-post", "post.md")
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("%s should not be copied to the built result", path)
	}
}

func TestBuildBlogIndexPage(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_generate_index")
	output := createTmpFolder(t)

	NewBlogBuilder(getConfiguration(), testDataDir).Build(output)

	bytes, _ := ioutil.ReadFile(filepath.Join(output, "index.html"))
	content := string(bytes)

	if !strings.Contains(content, `<meta http-equiv="X-UA-Compatible" content="IE=edge">`) {
		t.Fatalf("No base template in blog index file")
	}
	if !strings.Contains(content, "<a href=\"#\">Test Post</a>") {
		t.Fatalf("No post in blog index file")
	}
}

func TestAssetsWillBeCopied(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_copy_assets")
	output := createTmpFolder(t)

	err := NewBlogBuilder(getConfiguration(), testDataDir).Build(output)
	if err != nil {
		t.Fatal(err)
	}

	paths := []string{
		filepath.Join(output, "test.txt"),
		filepath.Join(output, "subdir", "test.txt"),
	}
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Fatalf("%s should exist in build result but not", path)
		}
	}
}

func TestGeneratedPostsWillBeSortedByDateInBlogIndex(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_posts_sorted_by_date")
	output := createTmpFolder(t)

	err := NewBlogBuilder(getConfiguration(), testDataDir).Build(output)
	if err != nil {
		t.Fatal(err)
	}

	bytes, _ := ioutil.ReadFile(filepath.Join(output, "index.html"))
	content := string(bytes)

	expectExcerpt := `
  <li><a href="2-second-post">2-second-post</a></li>

  <li><a href="1-first-post">1-first-post</a></li>
`
	if !strings.Contains(content, expectExcerpt) {
		t.Fatalf("Posts in blog index page should be sorted by date in descending")
	}
}
