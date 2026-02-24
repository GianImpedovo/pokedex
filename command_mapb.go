package main

import (
	"errors"
	"fmt"
)

func commandMapb(c *config) error {
	if c.Previous == nil {
		return errors.New("you're on the first page")
	}
	info, err := c.pokeapiClient.ListLocationAreas(c.Previous)
	if err != nil {
		return err
	}

	c.Next = info.Next
	c.Previous = info.Previous

	for _, v := range info.Results {
		fmt.Println(v.Name)
	}

	return nil

}
