package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Create a PostCreator that operate in specified directory
func NewPostCreator(dir string) PostCreator {
	return PostCreator{dir}
}

// A PostCreator create folder and files for a new post
type PostCreator struct {
	// directory to create post
	dir string
}

// Create a post with specified title
func (pc PostCreator) Create(title string) error {
	date := time.Now().Format(ISO8601Date)

	postDir := filepath.Join(pc.dir, Prettify(title))
	if err := os.Mkdir(postDir, os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s due to %v", postDir, err)
	}

	file, err := os.Create(filepath.Join(postDir, "post.md"))
	if err != nil {
		return fmt.Errorf("Unable to create post.md due to %v", err)
	}
	content := fmt.Sprintf(`---
date: %s
title: %s
---
`, date, title)
	file.WriteString(content)
	return nil
}

// PostParser parse file into Post
type PostParser struct {
	converter MarkdownConverter
	conf      Configuration
}

func NewPostParser(conf Configuration) PostParser {
	return PostParser{NewGithubMarkdownConverter(conf), conf}
}

// Parse will parse Post from specified post folder
func (pp PostParser) Parse(dir string) Post {
	post := NewPost()
	post.dir = dir
	post.key = filepath.Base(dir)
	post.variables, post.content = pp.parsePostFile(dir)
	post.htmlContent, _ = pp.converter.Convert(post.content)
	post.url = pp.conf.BaseUrl + filepath.Base(dir) + "/"
	return *post
}

func (pp PostParser) parsePostFile(dir string) (map[string]string, string) {

	f, _ := os.Open(filepath.Join(dir, PostFileName))
	defer f.Close()

	variables := make(map[string]string)
	content := ""

	isInVariablesBlock := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if !isInVariablesBlock && line == "---" {
			isInVariablesBlock = true
			continue
		} else if isInVariablesBlock && line == "---" {
			isInVariablesBlock = false
			continue
		}

		if isInVariablesBlock {
			if key, value, err := pp.parseVariable(line); err == nil {
				variables[key] = value
			}
		} else {
			content += (line + "\n")
		}
	}
	return variables, content
}

func (pp PostParser) parseVariable(line string) (key string, value string, err error) {
	pair := strings.SplitN(line, ":", 2)
	if len(pair) == 2 {
		key = strings.TrimSpace(pair[0])
		value = strings.TrimSpace(pair[1])
	} else {
		err = fmt.Errorf("Invalid post variable format : %s", line)
	}
	return
}

func NewPost() *Post {
	p := new(Post)
	return p
}

// Post represented a post in blog
type Post struct {
	// post folder name as the post identifier
	key string
	// variables in post file
	variables map[string]string
	// path of original post folder
	dir string
	// url of the post
	url string
	// content in Markdown
	content string
	// content in HTML converted from Markdown
	htmlContent string
}

func (p Post) Title() string {
	return p.variables["title"]
}

func (p Post) Date() string {
	return p.variables["date"]
}

func (p Post) RSSDate() string {
	return p.Time().Format(time.RFC1123)
}

func (p Post) Time() time.Time {
	t, err := time.Parse("2006-01-02", p.Date())
	if err != nil {
		panic(fmt.Sprintf("Post %s has an invalid date", p.Title()))
	}
	return t
}

func (p Post) Content() string {
	return p.content
}

func (p Post) Excerpt() string {
	if len(p.content) <= 200 {
		return p.content
	}
	return p.content[0:200] + "[...]"
}

func (p Post) HTMLContent() template.HTML {
	return template.HTML(p.htmlContent)
}

func (p Post) Key() string {
	return p.key
}

func (p Post) Dir() string {
	return p.dir
}

func (p Post) Url() string {
	return p.url
}

func (p Post) Variable(key string) string {
	return p.variables[key]
}

// Posts represented an array of Post
type Posts []Post

func (p Posts) Len() int {
	return len(p)
}

func (p Posts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type PostsByDateDesc struct {
	Posts
}

func (s PostsByDateDesc) Less(i, j int) bool {
	return !s.Posts[i].Time().Before(s.Posts[j].Time())
}
