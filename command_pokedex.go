package main

import (
	"fmt"
)

func commandPokedex(cfg *config, s string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Printf("No Pokemon Caught Yet...\nTry Catching One and Try Again\n")
		return nil
	} else {
		fmt.Printf("Your Pokedex:\n")
		for _, pokemon := range cfg.Pokedex {
			fmt.Printf("  - %s\n", pokemon.Name)
		}
	}
	return nil
}
