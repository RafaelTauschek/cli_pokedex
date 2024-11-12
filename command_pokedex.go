package main

import (
	"fmt"
)

func callbackPokedex(c *Client, params string) error {
	fmt.Println("Your Pokedex:")

	for _, pokemon := range c.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
