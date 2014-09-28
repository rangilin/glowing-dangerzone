package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// MarkdownConverter convert the specified markdown into HTML.
type MarkdownConverter interface {
	Convert(markdown string) string
}

// GithubMarkdownConverter convert markdown into HTML with Github's markdown API.
type GithubMarkdownConverter struct {
	conf Configuration
}

// NewGithubMarkdownConverter create a GithubMarkdownConverter with configuration
func NewGithubMarkdownConverter() GithubMarkdownConverter {
	gmc := GithubMarkdownConverter{}
	gmc.conf = getConfiguration()
	return gmc
}

// Convert convert markdown into HTML with Github api
func (gmc GithubMarkdownConverter) Convert(markdown string) (string, error) {
	request, err := http.NewRequest("POST",
		"https://api.github.com/markdown/raw", strings.NewReader(markdown))
	if err != nil {
		return "", err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "text/plain")
	request.Header.Add("Authorization", "token "+gmc.conf.GithubAccessToken)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Response is not okay, status code is %d",
			response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	return string(result), nil
}
