package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestPostParserParseVariable(t *testing.T) {
	file := withPostFileLikeThis(`---
key: value
---
`)

	post := new(GithubPostParser).Parse(file)

	assertPostHaveVariable(t, post, "key", "value")
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
