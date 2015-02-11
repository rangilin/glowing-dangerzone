package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCreatePost(t *testing.T) {
	t.Parallel()

	title := "Folder's name"
	date := time.Now().Format(ISO8601Date)
	dir := createTmpFolder(t)

	NewPostCreator(dir).Create(title)

	postDir := filepath.Join(dir, "folders-name")
	postFile := filepath.Join(postDir, "post.md")

	if _, err := os.Stat(postDir); os.IsNotExist(err) {
		t.Fatalf("Post folder %s should be created, but not", postDir)
	}
	if _, err := os.Stat(postFile); os.IsNotExist(err) {
		t.Fatalf("Post file %s should be created, but not", postFile)
	}

	expectContent := fmt.Sprintf(`---
date: %s
title: %s
---
`, date, title)
	bytes, _ := ioutil.ReadFile(postFile)
	actualContent := string(bytes)
	if actualContent != expectContent {
		t.Fatalf("Generated post file content is wrong, expect %s, but got %s ", expectContent, actualContent)
	}
}

func TestCreatePostWithDuplicatedTitle(t *testing.T) {
	t.Parallel()

	title := "duplicated"
	dir := createTmpFolder(t)
	postDir := filepath.Join(dir, "duplicated")
	os.Mkdir(postDir, os.ModePerm)

	err := NewPostCreator(dir).Create(title)

	if err == nil {
		t.Fatalf("Should error when create duplicated folder, but not")
	}
}

func TestParsePost(t *testing.T) {
	t.Parallel()

	testPostDir := "testdata/post_test/test-post-parser/"
	post := newTestPostParser().Parse(testPostDir)

	title := "This is a test"
	if post.Title() != title {
		t.Fatalf("Expect post title [%s], but got [%s]", title, post.Title())
	}

	date := "2015-01-08"
	if post.Date() != date {
		t.Fatalf("Expect post date [%s], but got [%s]", date, post.Date())
	}

	if post.Folder() != testPostDir {
		t.Fatalf("Expect post folder to be [%s], but got [%s]", testPostDir, post.Folder())
	}

	content := "content\n"
	if post.Content() != content {
		t.Fatalf("Expect post content [%s]. but got [%s]", content, post.Content())
	}

	htmlContent := "<html>content\n</html>"
	if post.HtmlContent() != htmlContent {
		t.Fatalf("Expect post html content [%s], but got [%s]", htmlContent,
			post.HtmlContent())
	}

	key := "test-post-parser"
	if post.Key() != key {
		t.Fatalf("Expect post key [%s], but got [%s]", key, post.Key())
	}

}

func newTestPostParser() PostParser {
	return PostParser{
		new(TestMarkdownConverter),
	}
}

// TestMarkdownConverter is a markdown converter that simply put markdown inside
// <html></html> tag, we use this to test without calling Github API
type TestMarkdownConverter struct {
}

// Convert simply return specified markdown.
func (tmdc TestMarkdownConverter) Convert(markdown string) (string, error) {
	return fmt.Sprintf("<html>%s</html>", markdown), nil
}
