package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/crbroughton/go-backstop/utils"
)

type Viewport struct {
	Name   string
	Width  int
	Height int
}

type Config struct {
	Viewports []Viewport
}

var path = "config.json"

func defaultViewports() Config {
	config := Config{
		Viewports: []Viewport{
			{
				Name:   "desktop",
				Width:  1280,
				Height: 720,
			},
			{
				Name:   "iPhone 12/13 Pro",
				Width:  390,
				Height: 844,
			},
		},
	}
	return config
}

func CreateJSON() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)

		if utils.IsError(err) {
			return
		}
		defer file.Close()
	}
}

func WriteDefaultConfiguration() {
	defaultJSON := defaultViewports()

	JSON, err := json.MarshalIndent(defaultJSON, "", " ")

	if utils.IsError(err) {
		log.Fatal("Failed to create default JSON configuration file")
	}

	err = os.WriteFile(path, JSON, 0644)

	if utils.IsError(err) {
		log.Fatal("Failed to write default configuration to the JSON file")
	}
}
