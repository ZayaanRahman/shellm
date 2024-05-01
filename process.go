package main

import (
	"fmt"
	"os"
)

func process(args []string) {

	switch parse(args) {

	case "ask":
		ask(args[1:])
	case "help":
		help(args[1:])
	case "model":
		model(args[1:])
	case "key":
		model(args[1:])
	case "prompt":
		model(args[1:])
	}
}

func parse(args []string) string {

	// default to help
	if len(args) == 0 || args[0] == "help" || args[0] == "h" {
		return "help"

	} else if args[0] == "ask" || args[0] == "a" {
		return "ask"

	} else if args[0] == "model" || args[0] == "m" {
		return "model"

	} else if args[0] == "key" || args[0] == "k" {
		return "key"

	} else if args[0] == "prompt" || args[0] == "p" {
		return "prompt"

	} else {
		fmt.Println("Invalid command entered.")
		os.Exit(0)
		return ""
	}
}
