package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/GianImpedovo/pokedexcli/internal/pokecache"
)

type Client struct {
	pokeapiClient http.Client
	pokeapiCache  pokecache.Cache
}

func NewClient(timeout, interval time.Duration) Client {

	return Client{
		pokeapiCache: pokecache.NewCache(interval),
		pokeapiClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) ListLocationAreas(url *string) (LocationAreaResponse, error) {
	fullUrl := baseURL + "/location-area"
	if url != nil {
		fullUrl = *url
	}

	data, ok := c.pokeapiCache.Get(fullUrl)
	var locations LocationAreaResponse
	if ok {
		// 2. Si está, la usamos
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locations, nil

	}
	// 3. Si no está, vamos a internet
	res, err := c.pokeapiClient.Get(fullUrl)
	if err != nil {
		return locations, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locations, err
	}

	c.pokeapiCache.Add(fullUrl, body)

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return locations, err
	}
	return locations, nil
}

func (c *Client) ListPokemonsArea(location string) (LocationPokemonsResponse, error) {
	fullurl := baseURL + "/location-area/" + location

	data, ok := c.pokeapiCache.Get(fullurl)
	var locations LocationPokemonsResponse
	if ok {
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return LocationPokemonsResponse{}, err
		}
		return locations, nil
	}

	res, err := c.pokeapiClient.Get(fullurl)
	if err != nil {
		return LocationPokemonsResponse{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationPokemonsResponse{}, err
	}

	c.pokeapiCache.Add(fullurl, body)

	var locationPokemons LocationPokemonsResponse
	err = json.Unmarshal(body, &locationPokemons)
	if err != nil {
		return LocationPokemonsResponse{}, err
	}

	return locationPokemons, nil
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	fullurl := baseURL + "/pokemon/" + pokemonName
	res, err := c.pokeapiClient.Get(fullurl)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	var pokemon Pokemon
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
