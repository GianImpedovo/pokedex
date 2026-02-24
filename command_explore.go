package main

import "fmt"

func commandExplore(c *config) error {
	fmt.Printf("Exploring %s...\n", c.cmd[1])
	fmt.Println("Found Pokemon:")

	list, err := c.pokeapiClient.ListPokemonsArea(c.cmd[1])
	if err != nil {
		return err
	}
	for _, p := range list.Pokemons {
		fmt.Printf("- %s\n", p.Pokemon.Name)
	}
	return nil
}
