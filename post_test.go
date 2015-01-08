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
	title := "Folder's name"
	date := time.Now().Format(ISO8601Date)
	dir := createTmpFolder(t)

	NewPostCreator(dir).Create(title)

	postDir := filepath.Join(dir, date+"-folders-name")
	postFile := filepath.Join(postDir, "post.md")

	assertFilePathExist(t, postDir)
	assertFilePathExist(t, postFile)
	assertFileContent(t, postFile, fmt.Sprintf(`---
date: %s
title: %s
---
`, date, title))
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

func TestPost(t *testing.T) {
	file := withPostFileLikeThis(`---
title: Post Title
---
test`)
	post := newTestPostParser().Parse(file)

	if post.Title() != "Post Title" {
		t.Fatalf("Expect post title is %s, but got %s", "Post Title", post.Title())
	}
}

func withPostFileLikeThis(content string) *os.File {
	file, _ := ioutil.TempFile(os.TempDir(), "post_")
	ioutil.WriteFile(file.Name(), []byte(content), 0644)
	return file
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
