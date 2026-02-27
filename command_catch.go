package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(cfg.cmd) < 2 {
		return fmt.Errorf("usage: catch <pokemon>")
	}
	pokemonName := cfg.cmd[1]
	info, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", info.Name)

	chance := 50 - info.BaseExperience/5
	if chance < 5 {
		chance = 5
	}
	if chance > 95 {
		chance = 95
	}
	roll := r.Intn(100)
	caught := roll < chance

	if caught {
		fmt.Printf("%s was caught!\n", info.Name)
		cfg.pokedex[info.Name] = info
	} else {
		fmt.Printf("%s escaped!\n", info.Name)
	}

	return nil
}
