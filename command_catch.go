package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func commandCatch(cfg *config, pokemonName string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName + "/"

	data, ok := cfg.Cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Error Calling %s Pokemon Endpoint: %w", pokemonName, err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Error Reading Return from %s Pokemon Endpoint: %w", pokemonName, err)
		}

		cfg.Cache.Add(url, data)
	}
	//Unmarshal JSON Response
	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		fmt.Printf("Not a valid Pokemon\nTry Again\n")
		return err
	}

	name := pokemon.Name
	baseExperience := pokemon.BaseExperience

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := r.Intn(baseExperience)
	caught := randomNum < 20

	if caught {
		fmt.Printf("%s was caught!\n", name)
		cfg.Pokedex[name] = pokemon
		return nil
	} else {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

}
