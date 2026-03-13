package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config, s string) error {
	var url string
	if cfg.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *cfg.Next
	}

	data, ok := cfg.Cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Error Calling Location-Area API Endpoint: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Error Reading Return from Location-Area API Endpoint: %w", err)
		}

		cfg.Cache.Add(url, data)
	}

	var respShallowLocation RespShallowLocations
	if err := json.Unmarshal(data, &respShallowLocation); err != nil {
		return err
	}
	//Print Area Names
	for _, area := range respShallowLocation.Results {
		fmt.Println(area.Name)
	}

	//Update cfg for Next and Previous URLs
	cfg.Next = respShallowLocation.Next
	cfg.Previous = respShallowLocation.Previous

	return nil
}

func commandMapB(cfg *config, s string) error {
	var url string
	if cfg.Previous == nil {
		fmt.Println("You Are on the First Page of Location-Areas\nTry Using map")
		return nil
	} else {
		url = *cfg.Previous
	}

	data, ok := cfg.Cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Error Calling Location-Area API Endpoint: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Error Reading Return from Locaiton-Area API Endpoint: %w", err)
		}

		cfg.Cache.Add(url, data)
	}

	var respShallowLocation RespShallowLocations
	if err := json.Unmarshal(data, &respShallowLocation); err != nil {
		return err
	}
	//Print Area Names
	for _, area := range respShallowLocation.Results {
		fmt.Println(area.Name)
	}

	//Update cfg for Next and Previous URLs
	cfg.Next = respShallowLocation.Next
	cfg.Previous = respShallowLocation.Previous

	return nil
}
