package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/crbroughton/go-backstop/utils"
)

type Viewport struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Scenario struct {
	Label string
	Url   string
}

type Config struct {
	Viewports []Viewport `json:"viewports"`
	Scenarios []Scenario `json:"scenarios"`
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
		Scenarios: []Scenario{},
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

		WriteDefaultConfiguration()

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

func AppendToViewportArray(newViewport Viewport) {
	// Read JSON file
	file, err := ioutil.ReadFile(path)
	if utils.IsError(err) {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal JSON
	var config Config
	err = json.Unmarshal(file, &config)
	if utils.IsError(err) {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Append the new Viewport to the Viewports array
	config.Viewports = append(config.Viewports, newViewport)

	// Marshal the updated Config struct back to JSON
	updatedJSON, err := json.MarshalIndent(config, "", "  ")
	if utils.IsError(err) {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write the updated JSON to a file
	err = ioutil.WriteFile(path, updatedJSON, 0644)
	if utils.IsError(err) {
		fmt.Println("Error writing file:", err)
		return
	}
}
