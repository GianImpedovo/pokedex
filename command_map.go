package main

import (
	"fmt"
)

func commandMap(c *config) error {

	info, err := c.pokeapiClient.ListLocationAreas(c.Next)
	if err != nil {
		return fmt.Errorf("could not list location areas: %w", err)
	}

	c.Next = info.Next
	c.Previous = info.Previous

	for _, v := range info.Results {
		fmt.Println(v.Name)
	}

	return nil

}
