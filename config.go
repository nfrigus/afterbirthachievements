package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// configType represents the configurable information of the application.
type configType struct {
	SteamAPIKey   string `json:"steamApiKey"`
	WebserverPort int    `json:"webserverPort"`
}

// Global variable to store the current configuration
var config configType

// readConfig reads the current configuration if it has not already been read.
func readConfig() (err error) {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		return errors.New("no config file found. Create one as config.json")
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return errors.New("Could not read config file.")
	}
	return
}
