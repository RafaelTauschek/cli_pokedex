package main

import (
	"fmt"
)

func callbackInspect(c *Client, params string) error {
	if params == "" {
		fmt.Println("You must provide a location area name.")
		return nil
	}

	pokemon, ok := c.caughtPokemon[params]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")

	for _, types := range pokemon.Types {
		fmt.Printf("- %s\n", types.Type.Name)
	}

	return nil
}
