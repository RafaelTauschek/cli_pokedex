package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/RafaelTauschek/pokedexcli/internal/pokecache"
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

		var params = ""
		if len(cleaned) > 1 {
			params = cleaned[1]
		}

		availableCommands := getComamnds()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		command.callback(c, params)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *Client, params string) error
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
		"explore": {
			name:        "explore",
			description: "Displayes all encounterable Pokenmons, write explore + location",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "Takes the name of a Pokemon as Argument, tries to catch it",
			callback:    callbackCatch,
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
	cache           *pokecache.Cache
	caughtPokemon   map[string]Pokemon
}

func NewClient() *Client {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	return &Client{
		client:          &http.Client{},
		nextLocationURL: &baseURL,
		prevLocationURL: nil,
		cache:           pokecache.NewCache(5 * time.Minute),
	}
}

func main() {
	client := NewClient()
	startRepl(client)
}
