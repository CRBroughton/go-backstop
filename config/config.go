package config

import (
	"encoding/json"
	"errors"
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
	Label string `json:"label"`
	Url   string `json:"url"`
}

type Cookie struct {
	Domain string `json:"domain"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

type Config struct {
	Viewports []Viewport `json:"viewports"`
	Scenarios []Scenario `json:"scenarios"`
	Cookies   []Cookie   `json:"cookies"`
}

const settingsFolder = ".settings"
const settingsPath = ".settings/config.json"

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
		Cookies:   []Cookie{},
	}
	return config
}

func createSettingsFolder() {
	var _, err = os.Stat(settingsFolder)

	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(settingsFolder, os.ModePerm)

		if utils.IsError(err) {
			log.Fatal("Failed to make settings folder")
		}
	}
}

func CreateJSON() {
	createSettingsFolder()

	var _, err = os.Stat(settingsPath)

	if os.IsNotExist(err) {
		var file, err = os.Create(settingsPath)

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

	err = os.WriteFile(settingsPath, JSON, 0644)

	if utils.IsError(err) {
		log.Fatal("Failed to write default configuration to the JSON file")
	}
}

// AppendToJSONArray appends a new struct to the configuration file
func AppendToJSONArray(newItem interface{}, fieldName string) {
	// Read JSON file
	file, err := ioutil.ReadFile(settingsPath)
	if utils.IsError(err) {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal JSON
	var data map[string]interface{}
	if err := json.Unmarshal(file, &data); err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Check if the field exists
	arr, ok := data[fieldName].([]interface{})
	if !ok {
		return
	}

	arr = append(arr, newItem)
	data[fieldName] = arr

	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if utils.IsError(err) {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write the updated JSON to a file
	err = ioutil.WriteFile(settingsPath, updatedJSON, 0644)
	if utils.IsError(err) {
		fmt.Println("Error writing file:", err)
		return
	}
}

func RunBackstopCommand(command string, withConfig bool) {
	workingDIR, err := os.Getwd()
	if utils.IsError(err) {
		log.Fatal(err)
	}

	var JSConfig string

	if withConfig {
		JSConfig = JSConfig + "--config=.settings/backstop.config.js"
	}

	args := []string{
		"run",
		"--rm",
		"-v",
		workingDIR + ":/src",
		"backstopjs/backstopjs",
		command,
		JSConfig,
	}
	_, err = utils.RunCommand("docker", args...)

	if utils.IsError(err) {
		log.Fatal(err)
	}
}
