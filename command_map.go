package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Result   []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config) error {
	url := c.Next
	info, ok := c.pokeapiClient.Get(url)
	var data locationAreaResponse
	if ok {
		err := json.Unmarshal(info, &data)
		if err != nil {
			return fmt.Errorf("could not parse response body: %w", err)
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("could not get location area response: %w", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		c.pokeapiClient.Add(url, body)
		err = json.Unmarshal(body, &data)
		if err != nil {
			return fmt.Errorf("could not read response body: %w", err)
		}

	}

	c.Next = data.Next
	c.Previous = data.Previous

	for _, v := range data.Result {
		fmt.Println(v.Name)
	}

	return nil

}
