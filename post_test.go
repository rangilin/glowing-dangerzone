package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCreatePostFolder(t *testing.T) {
	title := "Folder's name"
	dir := createTmpFolder(t)

	NewPostCreator(dir).Create(title)

	postDir := time.Now().Format(ISO8601Date) + "-folders-name"
	assertFilePathExist(t, filepath.Join(dir, postDir))
}

func TestCreatePostFile(t *testing.T) {
	title := "test"
	dir := createTmpFolder(t)

	NewPostCreator(dir).Create(title)

	postDir := time.Now().Format(ISO8601Date) + "-" + title
	assertFilePathExist(t, filepath.Join(dir, postDir, "post.md"))
}

func TestPostFileContent(t *testing.T) {
	title := "Post Title"
	dir := createTmpFolder(t)

	NewPostCreator(dir).Create(title)

	postDir := time.Now().Format(ISO8601Date) + "-post-title"
	postFilePath := filepath.Join(dir, postDir, "post.md")
	content, _ := ioutil.ReadFile(postFilePath)

	expect := `---
title: Post Title
---
`
	if expect != string(content) {
		t.Fatalf("Expect post file be \n%s\n, but got \n%s\n", expect, content)
	}
}

func TestParseVariable(t *testing.T) {
	file := withPostFileLikeThis(`---
key1: value1
key2: value2
---
`)

	post := newTestPostParser().Parse(file)

	assertPostHaveVariable(t, post, "key1", "value1")
	assertPostHaveVariable(t, post, "key2", "value2")
}

func TestParseContent(t *testing.T) {
	file := withPostFileLikeThis(`---
---
test`)
	post := newTestPostParser().Parse(file)

	assertPostContent(t, post, "test\n")
}

func TestParseHtmlContent(t *testing.T) {
	file := withPostFileLikeThis(`---
---
test`)
	post := newTestPostParser().Parse(file)

	assertPostHtmlContent(t, post, "<html>test\n</html>")
}

func withPostFileLikeThis(content string) os.File {
	file, _ := ioutil.TempFile(os.TempDir(), "post_")
	ioutil.WriteFile(file.Name(), []byte(content), 0644)
	return *file
}

func assertPostHaveVariable(t *testing.T, post Post, key string, expectedValue string) {
	if value, exist := post.variables[key]; !exist || value != expectedValue {
		t.Fatalf("Expect a variable pair [%s: %s] exist in post, but not exist",
			key, expectedValue)
	}
}

func assertPostContent(t *testing.T, post Post, expected string) {
	if post.content != expected {
		t.Fatalf("Expect post content is %q, but got %q", expected, post.content)
	}
}

func assertPostHtmlContent(t *testing.T, post Post, expected string) {
	if post.htmlContent != expected {
		t.Fatalf("Expect post html content is %q, but got %q", expected, post.htmlContent)
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
func (tmdc TestMarkdownConverter) Convert(markdown string) string {
	return fmt.Sprintf("<html>%s</html>", markdown)
}
