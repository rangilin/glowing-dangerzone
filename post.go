package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PostParser parse file into Post
type PostParser struct {
	converter MarkdownConverter
}

// Parse parse specified file into post
func (pp PostParser) Parse(f os.File) Post {
	post := NewPost()
	post.variables, post.content = parseLineByLine(f)
	post.htmlContent = pp.converter.Convert(post.content)
	return *post
}

func parseLineByLine(f os.File) (variables map[string]string, content string) {
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
			if key, value, err := parseVariable(line); err == nil {
				variables[key] = value
			}
		} else {
			content += (line + "\n")
		}
	}

	return
}

func parseVariable(line string) (key string, value string, err error) {
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
