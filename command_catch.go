package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type Pokemon struct {
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func callbackCatch(c *Client, params string) error {
	if params == "" {
		fmt.Println("You didn't provide a Pokemon to catch")
	}

	pokemonURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", params)

	req, err := http.Get(pokemonURL)
	if err != nil {
		return err
	}

	res, err := io.ReadAll(req.Body)
	req.Body.Close()

	if err != nil {
		return err
	}

	var pokemon Pokemon

	err = json.Unmarshal(res, &pokemon)
	if err != nil {
		return err
	}

	normalized := 1 - (float64(pokemon.BaseExperience-40) / float64(300))
	catchRate := (normalized * 60) + 20
	roll := rand.Intn(100)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if roll > int(catchRate) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		c.caughtPokemon[pokemon.Name] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
