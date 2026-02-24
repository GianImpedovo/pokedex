package main

import "fmt"

func commandMapb(c *config) error {
	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return fmt.Errorf("you're on the first page")
	}
	info, err := c.pokeapiClient.ListLocationAreas(c.Previous)
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
