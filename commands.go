package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Body struct {
	Model    string              `json:"model"`
	Messages []map[string]string `json:"messages"`
}

type Response struct {
	Error *struct {
		Message string      `json:"message"`
		Type    string      `json:"type"`
		Param   interface{} `json:"param"`
		Code    string      `json:"code"`
	} `json:"error"`

	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func help(args []string) {
	fmt.Printf(`

____       _          _     _     __  __ 
\ \ \  ___| |__   ___| |   | |   |  \/  |
 \ \ \/ __| '_ \ / _ \ |   | |   | |\/| |
 / / /\__ \ | | |  __/ |___| |___| |  | |
/_/_/ |___/_| |_|\___|_____|_____|_|  |_|
	
_____ _____ _____ _____ _____ _____ _____
|_____|_____|_____|_____|_____|_____|_____|

> Welcome to shellm!

Enter "shellm ask [your query]" to generate a response.
For other options/commands, enter "shellm [command name] [params, if any]".

Commands:
	help    OR h   View available commands.
	ask     OR a   Enter a query. All input after the "ask" command will be interpreted as part of the query.
	model   OR m   View the current model and select a model from the list (not implemented). Default gpt-3.5-turbo.
	key     OR k   View the current key, or enter one as a parameter. A key must be provided before use.
	prompt  OR p   View the current system prompt, or enter a new one as a parameter.
	
This is a WIP! All commands might not be fully implemented yet.
_____ _____ _____ _____ _____ _____ _____  
|_____|_____|_____|_____|_____|_____|_____|
`)
}

func ask(args []string) {

	query := strings.Join(args, " ")
	config := getConfig()
	if config.Key == "" {
		fmt.Println("No key provided")
		os.Exit(0) // replace with better error handling
	}

	body := Body{Model: config.Model, Messages: []map[string]string{

		{"role": "system", "content": config.Prompt},
		{"role": "user", "content": query},
	}}

	reqData, err := json.Marshal(body)

	if err != nil {
		fmt.Println(err.Error())
	}

	endpoint := "https://api.openai.com/v1/chat/completions"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqData))

	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	response := Response{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println(err.Error())
	}

	if response.Error != nil {
		fmt.Printf("Error calling API; error code: %s\n", response.Error.Code)

	} else {

		fmt.Println(response.Choices[0].Message.Content)
	}
}

func model(args []string) {
	fmt.Printf("Current model: gpt-3.5-turbo\n")
}

func key(args []string) {

	config := getConfig()

	if len(args) == 0 {
		fmt.Printf("Current key: %s\n", config.Key)

	} else if len(args) == 1 {
		oldKey := config.Key
		config.Key = args[0]
		updateConfig(config)
		fmt.Printf("Old key: %s\nNew key: %s\n", oldKey, config.Key)

	} else {
		fmt.Println("Invalid number of args")
	}
}

func prompt(args []string) {

	config := getConfig()

	if len(args) == 0 {
		fmt.Printf("Current prompt: %s\n", config.Prompt)

	} else {
		newPrompt := strings.Join(args, " ")

		oldPrompt := config.Prompt
		config.Prompt = newPrompt
		updateConfig(config)
		fmt.Printf("\nOld prompt: %s\n\nNew prompt: %s\n", oldPrompt, newPrompt)
	}
}
