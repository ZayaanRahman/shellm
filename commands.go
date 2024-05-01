package main

import (
	"fmt"
)

func help(args []string) {
	fmt.Printf(`

     _          _     _     __  __ 
 ___| |__   ___| |   | |   |  \/  |
/ __| '_ \ / _ \ |   | |   | |\/| |
\__ \ | | |  __/ |___| |___| |  | |
|___/_| |_|\___|_____|_____|_|  |_|
	
_____ _____ _____ _____ _____ _____ _____
|_____|_____|_____|_____|_____|_____|_____|

Welcome to shellm!

Enter "shellm ask [your query]" to generrate a response
For other options/commands, enter "shellm [command name] [params, if any]"

Commands:
	help    OR k   View available commands.
	ask     OR a   Enter a query. All input after the "ask" command will be interpreted as part of the query.
	model   OR m   View the current model and select a model from the list (not implemented). Default gpt-3.5-turbo.
	key     OR k   View the current key, or enter one as a parameter. A key must be provided before use.
	prompt  OR p   View the current system prompt, or enter a new one as a parameter.
	
This is a WIP! All commands might not be fully implemented yet.
_____ _____ _____ _____ _____ _____ _____
|_____|_____|_____|_____|_____|_____|_____|`)
}

func ask(args []string) {
	fmt.Printf("ask selected")
}

func model(args []string) {
	fmt.Printf("model selected")
}

func key(args []string) {
	fmt.Printf("key selected")
}

func prompt(args []string) {
	fmt.Printf("prompt selected")
}
