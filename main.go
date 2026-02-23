package main

import (
	"time"

	"github.com/GianImpedovo/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient pokecache.Cache
	Next          string
	Previous      string
}

func main() {
	interval := 5 * time.Minute
	cache := pokecache.NewCache(interval)

	cfg := &config{
		pokeapiClient: *cache,
		Next:          "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		Previous:      "",
	}
	startRepl(cfg)
}
