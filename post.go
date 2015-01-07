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

	postDir := filepath.Join(pc.dir, date+"-"+title)
	if err := os.Mkdir(postDir, os.ModePerm); err != nil {
		return fmt.Errorf("Unable to create folder %s", postDir)
	}

	if _, err := os.Create(filepath.Join(postDir, "post.md")); err != nil {
		return fmt.Errorf("Unable to create post.md")
	}
	return nil
}

// PostParser parse file into Post
type PostParser struct {
	converter MarkdownConverter
}

// Parse parse specified file into post
func (pp PostParser) Parse(f os.File) Post {
	post := NewPost()
	post.variables, post.content = pp.parseLineByLine(f)
	post.htmlContent = pp.converter.Convert(post.content)
	return *post
}

func (pp PostParser) parseLineByLine(f os.File) (variables map[string]string, content string) {
	variables = make(map[string]string)
	content = ""

	defer f.Close()

	isInVariablesBlock := false
	scanner := bufio.NewScanner(&f)
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

	return
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

// Post represented a post in blog
type Post struct {
	variables   map[string]string
	content     string
	htmlContent string
}

func NewPost() *Post {
	p := new(Post)
	return p
}
