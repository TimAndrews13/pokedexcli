package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemonName string) error {
	pokemon, exists := cfg.Pokedex[pokemonName]

	if !exists {
		fmt.Print("you have not caught that pokemon\n")
		return nil
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stats := range pokemon.Stats {
			fmt.Printf("  - %s: %d\n", stats.Stat.Name, stats.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, types := range pokemon.Types {
			fmt.Printf("  - %s\n", types.Type.Name)
		}
	}
	return nil
}
