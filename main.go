package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func startRepl(c *Client) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getComamnds()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		command.callback(c)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *Client) error
}

func getComamnds() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Displays the name of the next 20 locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displayes the name of the previous 20 locations",
			callback:    callbackMapB,
		},
	}
}

func cleanInput(str string) []string {
	lowerd := strings.ToLower(str)
	words := strings.Fields(lowerd)
	return words
}

type Client struct {
	client          *http.Client
	nextLocationURL *string
	prevLocationURL *string
}

func NewClient() *Client {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	return &Client{
		client:          &http.Client{},
		nextLocationURL: &baseURL,
		prevLocationURL: nil,
	}
}

func main() {
	client := NewClient()
	startRepl(client)
}
