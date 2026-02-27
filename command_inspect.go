package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config) error {
	if len(cfg.cmd) < 2 {
		return errors.New("pokemon name is required")
	}

	pokemonName := cfg.cmd[1]
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	return nil
}
