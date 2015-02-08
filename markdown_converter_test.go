package main

import (
	"testing"
)

func TestGithubMarkdownConverter(t *testing.T) {
	t.Parallel()

	gmc := NewGithubMarkdownConverter(getConfiguration())

	html, err := gmc.Convert("Hello World")

	if err != nil {
		t.Fatalf("Unable to convert markdown : %s", err.Error())
	}

	expect := "<p>Hello World</p>\n"
	if html != expect {
		t.Fatalf("Expect html is %s, but got %q ", expect, html)
	}
}
