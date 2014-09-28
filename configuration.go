package main

import (
	"encoding/json"
	"os"
)

// Configuration of the blog engine
type Configuration struct {
	// Access token of github
	GithubAccessToken string
}

// Read configuration from config.json file
func getConfiguration() Configuration {
	file, err := os.Open("config.json")
	if err != nil {
		panic("Unable to open configuration file : " + err.Error())
	}

	decoder := json.NewDecoder(file)

	conf := Configuration{}
	if err := decoder.Decode(&conf); err != nil {
		panic("Unable to decode configuration : " + err.Error())
	}
	return conf
}
