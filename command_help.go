package main

import "fmt"

func callbackHelp(c *Client) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")

	availableCommands := getComamnds()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}
