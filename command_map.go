package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationsJSON struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []LocationsResultJSON
}

type LocationsResultJSON struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func callbackMap(c *Client) error {
	if c.nextLocationURL == nil {
		fmt.Println("You're already at the end!")
		return nil
	}

	res, err := http.Get(*c.nextLocationURL)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return err
	}

	var locations LocationsJSON
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}

	c.nextLocationURL = locations.Next
	c.prevLocationURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func callbackMapB(c *Client) error {
	if c.prevLocationURL == nil {
		fmt.Println("You're already at the start!")
		return nil
	}

	res, err := http.Get(*c.prevLocationURL)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	var locations LocationsJSON
	err = json.Unmarshal(body, &locations)

	if err != nil {
		return err
	}

	c.nextLocationURL = locations.Next
	c.prevLocationURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
