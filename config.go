package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Reflects fields in "config.json"
type Config struct {
	Model  string `json:"model"`
	Key    string `json:"key"`
	Prompt string `json:"prompt"`
	Os     string `json:"os"`
}

// Returns a config with default settings
func DefaultConfig() Config {
	return Config{
		Model:  "gpt-3.5-turbo",
		Key:    "",
		Prompt: "You are a ChatGPT assistant meant to help users navigate and learn their command line interface for the Windows operating system. Your goal is to provide the user with the most appropriate information about what commands to use in their terminal to accomplish some task, based on their query. Users will enter queries, usually asking about what command to use to do something with their operating system or use some command line tool. Try to limit your responses to around 50 characters, telling them succinctly which terminal command to use. Again, try to keep the responses short so they fit on one line. If the user asks about something unrelated to the terminal, you can give them a longer response that spans multiple lines. Or, if they want to do something complicated that would take multiple commands, you may use multiple lines to tell them how to accomplish what they want. Still, try to keep it short. Be brief, especially when asnwering questions about terminal commands. Prioritize correctness and brevity.",
		Os:     "Windows",
	}
}

// Get a configuration object reflecting the current "config.json" file
func getConfig() Config {

	file, err := os.Open("config.json")

	if err != nil {

		if os.IsNotExist(err) {
			fmt.Print("Creating config file...")
			updateConfig(DefaultConfig()) // create a default "config.json" if it doesn't yet exist
			fmt.Println("done")
		} else {
			fmt.Println(err.Error())
		}
	}
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err.Error())
	}

	curConfig := Config{}
	json.Unmarshal(data, &curConfig)
	return curConfig
}

// Update the "config.json" file with a new configuration object
func updateConfig(config Config) {

	data, err := json.Marshal(config)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.WriteFile("config.json", data, 0644)

	if err != nil {
		fmt.Println(err.Error())
	}
}
