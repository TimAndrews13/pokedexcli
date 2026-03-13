package main

import (
	"strings"

	"github.com/timandrews/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
	Pokedex  map[string]Pokemon
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
		"map": {
			name:        "map",
			description: "Displays Next 20 location-areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays Previous 20 location-areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explores Location-Area and Displays Pokemon from Location-Area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to Catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Provides Info and Stats for Caught Pokemon",
			callback:    commandInspect,
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
