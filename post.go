package main

import (
	"bufio"
	"fmt"
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

	postDir := filepath.Join(pc.dir, date+"-"+Prettify(title))
	if err := os.Mkdir(postDir, os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", postDir)
	}

	file, err := os.Create(filepath.Join(postDir, "post.md"))
	if err != nil {
		return fmt.Errorf("Unable to create post.md")
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
}

func NewPostParser() PostParser {
	return PostParser{NewGithubMarkdownConverter()}
}

// Parse will parse Post from specified post folder
func (pp PostParser) Parse(dir string) Post {
	post := NewPost()

	post.key = filepath.Base(dir)
	post.variables, post.content = pp.parsePostFile(dir)
	post.htmlContent, _ = pp.converter.Convert(post.content)
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

func (p Post) Content() string {
	return p.content
}

func (p Post) HtmlContent() string {
	return p.htmlContent
}

func (p Post) Key() string {
	return p.key
}
