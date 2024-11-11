package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EncounterJSON struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func callbackExplore(c *Client, params string) error {

	if params == "" {
		fmt.Println("You must provide a location area name.")
	}

	if params != "" {
		fmt.Printf("Params recieved %s\n", params)
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
