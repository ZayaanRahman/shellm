package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

type Config struct {
	Model  string `json:"model"`
	Key    string `json:"key"`
	Prompt string `json:"prompt"`
}

func updateConfig(args []string) {

	var curConfig Config

	json.Unmarshal(readFile(), &curConfig)

	newConfig := Config{Model: curConfig.Model, Key: curConfig.Key, Prompt: curConfig.Prompt}

	for _, arg := range args {

		indexEq := strings.IndexRune(arg, '=')

		if utf8.RuneCountInString(arg) >= 2 && arg[0:2] != "--" {
			fmt.Println(`Options must begin with "--".`)

		} else if indexEq == -1 || indexEq == 0 || indexEq == utf8.RuneCountInString(arg)-1 {
			fmt.Println(`Must set option name equal "=" to new setting. For example:
gpterminal key=ABCDEF`)

		} else {

			option := arg[2:indexEq]
			newStr := arg[indexEq+1 : utf8.RuneCountInString(arg)]

			switch option {

			case "model":
				// temp until more models added
				if newStr != "gpt4" {
					fmt.Println("Model name is invalid")
				}

			case "key":

				newConfig.Key = newStr
			// case "prompt":

			// 	// temp until figure out better way to edit
			// 	fmt.Println("Model name is invalid")
			default:

				fmt.Println("Invalid option")
			}
		}
	}

	file, err := os.Create("config.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	newJson, err := json.Marshal(newConfig)

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = file.Write(newJson)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func readFile() []byte {

	file, err := os.Open("config.json")

	if err != nil {

		if os.IsNotExist(err) {
			fmt.Println("Config file could not be found.")
		} else {
			fmt.Println(err.Error())
		}
	}
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err.Error())
	}

	return data
}
