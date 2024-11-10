package main

import "os"

func callbackExit(c *Client) error {
	os.Exit(0)
	return nil
}
