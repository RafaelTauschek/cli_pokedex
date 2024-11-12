package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EncounterJSON struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func callbackExplore(c *Client, params string) error {

	if params == "" {
		fmt.Println("You must provide a location area name.")
	}

	locationURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", params)

	data, found := c.cache.Get(locationURL)

	if !found {
		req, err := http.Get(locationURL)
		if err != nil {
			return err
		}

		data, err = io.ReadAll(req.Body)
		req.Body.Close()

		if err != nil {
			return err
		}
		c.cache.Add(locationURL, data)
	}

	var encounter EncounterJSON
	err := json.Unmarshal(data, &encounter)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", params)
	fmt.Println("Found Pokemon:")
	for _, enc := range encounter.PokemonEncounters {
		fmt.Printf("- %s\n", enc.Pokemon.Name)
	}

	return nil
}
