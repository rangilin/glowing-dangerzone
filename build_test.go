package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreateBlogFolder(t *testing.T) {
	t.Parallel()

	blogFileDir := testDataPath("build", "test_generate_files")
	outputDir := filepath.Join(createTmpFolder(t), BuildDirName)

	err := NewBlogBuilder(newTestPostParser(), fakeConfiguration(), blogFileDir).Build(outputDir)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		t.Fatalf("A folder call 'blog' should be created at %s after build.", outputDir)
	}
}

func TestCleanUpBeforeBuild(t *testing.T) {
	t.Parallel()

	blogFileDir := testDataPath("build", "test_generate_files")
	outputDir := filepath.Join(createTmpFolder(t), BuildDirName)

	deleteme := filepath.Join(outputDir, "delete_me")
	os.Mkdir(outputDir, os.ModePerm)
	os.Create(deleteme)

	err := NewBlogBuilder(newTestPostParser(), fakeConfiguration(), blogFileDir).Build(outputDir)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(deleteme); !os.IsNotExist(err) {
		t.Fatalf("Exist blog folder should be clean up before build")
	}
}

func TestBuildPost(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_generate_posts")
	output := createTmpFolder(t)

	err := NewBlogBuilder(newTestPostParser(), fakeConfiguration(), testDataDir).Build(output)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := ioutil.ReadFile(filepath.Join(output, "test-post", "index.html"))
	if err != nil {
		t.Fatal(err)
	}

	content := string(bytes)
	if !strings.Contains(content, `<meta http-equiv="X-UA-Compatible" content="IE=edge">`) {
		t.Errorf("Post index file should be generated with base template")
	}

	if !strings.Contains(content, "This is test post content") {
		t.Errorf("Post data should be available in post index file")
	}

	if !strings.Contains(content, "GithubAccessToken") {
		t.Errorf("Configuration data should be available in post index file")
	}
}

func TestBuildShouldCopyPostFiles(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_generate_posts")
	output := createTmpFolder(t)

	err := NewBlogBuilder(newTestPostParser(), fakeConfiguration(), testDataDir).Build(output)
	if err != nil {
		t.Fatal(err)
	}

	file := filepath.Join(output, "test-post", "test.txt")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("File in post folder should be copied to build result")
	}

	fileInSubDir := filepath.Join(output, "test-post", "test", "test.txt")
	if _, err := os.Stat(fileInSubDir); os.IsNotExist(err) {
		t.Errorf("File in sub folder of post folder should be copied to the build result")
	}

	markdown := filepath.Join(output, "test-post", "post.md")
	if _, err := os.Stat(markdown); !os.IsNotExist(err) {
		t.Errorf("Post markdown file should not be copied to the built result")
	}

}

func TestBuildBlogIndexPage(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_generate_index")
	output := createTmpFolder(t)

	err := NewBlogBuilder(newTestPostParser(), fakeConfiguration(), testDataDir).Build(output)
	if err != nil {
		t.Fatal(err)
	}

	index := filepath.Join(output, "index.html")
	if _, err := os.Stat(index); os.IsNotExist(err) {
		t.Fatalf("Blog index file should be created at %s", index)
	}

	bytes, _ := ioutil.ReadFile(index)
	content := string(bytes)
	if !strings.Contains(content, `<meta http-equiv="X-UA-Compatible" content="IE=edge">`) {
		t.Errorf("Blog index file should be generated with base template")
	}
	if !strings.Contains(content, "<a href=\"#\">Test Post</a>") {
		t.Errorf("Blog index file should have posts data available")
	}
	if !strings.Contains(content, "GithubAccessToken") {
		t.Errorf("Blog index file should have configuration data available")
	}
}

func TestAssetsWillBeCopied(t *testing.T) {
	t.Parallel()

	testDataDir := testDataPath("build", "test_copy_assets")
	output := createTmpFolder(t)

	err := NewBlogBuilder(newTestPostParser(), fakeConfiguration(), testDataDir).Build(output)
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

	err := NewBlogBuilder(newTestPostParser(), fakeConfiguration(), testDataDir).Build(output)
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
