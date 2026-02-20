package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(c *config) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		c.Next = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
		return nil
	}
	res, err := http.Get(c.Previous)
	if err != nil {
		return fmt.Errorf("could not get location area response: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %w", err)
	}

	var data locationAreaResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return fmt.Errorf("could not parse response body: %w", err)
	}

	c.Next = data.Next
	c.Previous = data.Previous

	for _, v := range data.Result {
		fmt.Println(v.Name)
	}

	return nil

}
