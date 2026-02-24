package main

import (
	"time"

	"github.com/GianImpedovo/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	cmd           []string
}

func main() {
	interval := 5 * time.Minute
	pokeClient := pokeapi.NewClient(5*time.Second, interval)

	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
