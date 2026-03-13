package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(cfg *config, area string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + area + "/"

	data, ok := cfg.Cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Error Calling %s Location Area Endpoint: %w", area, err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Error Reading Return from %s Location-Area API Endpoint: %w", area, err)
		}

		cfg.Cache.Add(url, data)
	}
	//Unmarshal JSON Response
	var locationDetail LocationDetail
	if err := json.Unmarshal(data, &locationDetail); err != nil {
		fmt.Printf("Not a valid Location-Area\nTry Again\n")
		return err
	}

	fmt.Printf("Exploring %s... \n", area)

	for _, encounter := range locationDetail.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
