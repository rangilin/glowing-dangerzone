package main

import (
	"fmt"
	"html/template"
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
		t.Errorf("Expect post title [%s], but got [%s]", title, post.Title())
	}

	date := "2015-01-08"
	if post.Date() != date {
		t.Errorf("Expect post date [%s], but got [%s]", date, post.Date())
	}

	rssDate := "Thu, 08 Jan 2015 00:00:00 UTC"
	if post.RSSDate() != rssDate {
		t.Errorf("Expect post rss date [%s], but got [%s]", rssDate, post.RSSDate())
	}

	if post.Dir() != testPostDir {
		t.Errorf("Expect post folder to be [%s], but got [%s]", testPostDir, post.Dir())
	}

	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi lacinia tempor purus, vitae aliquam elit. Morbi efficitur ut ante vehicula vestibulum. Curabitur pellentesque quam non pulvinar dictum. Sed semper augue eu massa ultricies bibendum. Nunc eleifend, dolor aliquam auctor placerat, enim ante convallis tortor, non laoreet mi enim eget ex. Proin semper augue non lorem luctus blandit. Suspendisse potenti.\n"
	if post.Content() != content {
		t.Errorf("Expect post content [%s]. but got [%s]", content, post.Content())
	}

	htmlContent := template.HTML("<html>" + content + "</html>")
	if post.HTMLContent() != htmlContent {
		t.Errorf("Expect post html content [%s], but got [%s]", htmlContent,
			post.HTMLContent())
	}

	key := "test-post-parser"
	if post.Key() != key {
		t.Errorf("Expect post key [%s], but got [%s]", key, post.Key())
	}

	url := "http://localhost/test-post-parser/"
	if post.Url() != url {
		t.Errorf("Expect post url [%s], but got [%s]", url, post.Url())
	}

}

func TestPostExcerpt(t *testing.T) {
	post := newTestPostParser().Parse("testdata/post_test/long-content-post/")
	excerpt := `一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十[...]`
	if post.Excerpt() != excerpt {
		t.Errorf("Expect post excerpt [%s], but got [%s]", excerpt, post.Excerpt())
	}
}

func newTestPostParser() PostParser {
	return PostParser{
		new(TestMarkdownConverter),
		fakeConfiguration(),
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
