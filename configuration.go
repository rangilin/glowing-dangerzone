package main

import (
	"os"
)

// Configuration of the blog engine
type Configuration struct {
	// Access token of github
	GithubAccessToken string
	// URL of the blog, should end with a slash
	BaseUrl string
}

// Read configuration from config.json file
func getConfiguration() Configuration {
	return Configuration{
		GithubAccessToken: os.Getenv("GD_GITHUB_ACCESS_TOKEN"),
		BaseUrl:           os.Getenv("GD_BASE_URL"),
	}
}
