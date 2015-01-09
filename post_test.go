package main

import (
	"fmt"
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

func TestParsePost(t *testing.T) {
	testPostDir := "testdata/post_test/2015-01-08-test-post-parser/"
	post := newTestPostParser().Parse(testPostDir)
	title := "This is a test"
	if post.Title() != title {
		t.Fatalf("Expect post title [%s], but got [%s]", title, post.Title())
	}

	date := "2015-01-08"
	if post.Date() != date {
		t.Fatalf("Expect post date [%s], but got [%s]", date, post.Date())
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

	key := "2015-01-08-test-post-parser"
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
