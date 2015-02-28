package main

import (
	"os"
	"testing"
)

func TestConfigurationIsReadFromEnvironment(t *testing.T) {
	// Should not run in parallel since we change environment variable

	// restore environment later
	token := os.Getenv("GD_GITHUB_ACCESS_TOKEN")
	defer os.Setenv("GD_GITHUB_ACCESS_TOKEN", token)

	os.Setenv("GD_GITHUB_ACCESS_TOKEN", "whatever")

	conf := getConfiguration()

	if conf.GithubAccessToken != "whatever" {
		t.Fatalf("Configuration should read from environment, but not")
	}
}

func fakeConfiguration() Configuration {
	return Configuration{
		GithubAccessToken: "GithubAccessToken",
		BaseUrl:           "http://localhost/",
	}
}
