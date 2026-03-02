package main

import "fmt"

func commandPokedex(cfg *config) error {
	pokedex := cfg.pokedex

	fmt.Println("Your Pokedex")
	for _, v := range pokedex {
		fmt.Printf("- %s\n", v.Name)
	}
	return nil
}
