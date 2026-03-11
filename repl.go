package main

import (
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	return supportedCommands
}

func cleanInput(text string) []string {
	var result_string []string

	text = strings.ToLower(text)

	result_string = strings.Fields(text)

	return result_string
}
