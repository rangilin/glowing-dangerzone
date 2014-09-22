package main

import (
	"os"
)

// PostParser parse file into Post
type PostParser interface {

	// Parse specified file into post
	Parse(f os.File) Post
}

// Post represented a post in blog
type Post struct {
	variables map[string]string
}

// Post parser that call Github's API to generate html
type GithubPostParser struct {
}

func (pp GithubPostParser) Parse(f os.File) Post {
	return Post{}
}
