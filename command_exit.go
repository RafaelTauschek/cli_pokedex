package main

import "os"

func callbackExit(c *Client, params string) error {
	os.Exit(0)
	return nil
}
