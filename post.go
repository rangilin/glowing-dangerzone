package main

import (
	"bufio"
	"os"
	"strings"
)

// PostParser parse file into Post
type PostParser struct {
}

// Parse specified file into post
func (pp PostParser) Parse(f os.File) Post {
	defer f.Close()

	post := NewPost()
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
			pair := strings.SplitN(line, ":", 2)
			if len(pair) == 2 {
				key := strings.TrimSpace(pair[0])
				value := strings.TrimSpace(pair[1])
				post.variables[key] = value
			}

		}
	}
	return *post
}

// Post represented a post in blog
type Post struct {
	variables map[string]string
}

func NewPost() *Post {
	p := new(Post)
	p.variables = make(map[string]string)
	return p
}
