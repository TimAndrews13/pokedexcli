package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Print("\n")
	// loop through supportCommands dynamically
	commands := getCommands()
	for command, value := range commands {
		fmt.Printf("%s: %s\n", command, value.description)
	}

	return nil
}
