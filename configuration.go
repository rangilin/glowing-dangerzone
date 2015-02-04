package main

import (
	"os"
)

// Configuration of the blog engine
type Configuration struct {
	// Access token of github
	GithubAccessToken string
}

// Read configuration from config.json file
func getConfiguration() Configuration {
	return Configuration{
		GithubAccessToken: os.Getenv("GD_GITHUB_ACCESS_TOKEN"),
	}
}
