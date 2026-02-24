package pokeapi

import (
	"encoding/json"
	"fmt"
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
			return locations, fmt.Errorf("could not parse response body: %w", err)
		}

	} else {
		// 3. Si no está, vamos a internet
		res, err := c.pokeapiClient.Get(fullUrl)
		if err != nil {
			return locations, fmt.Errorf("could not get location area response: %w", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return locations, fmt.Errorf("could not read response body: %w", err)
		}

		c.pokeapiCache.Add(fullUrl, body)

		err = json.Unmarshal(body, &locations)
		if err != nil {
			return locations, fmt.Errorf("could not parse response body: %w", err)
		}
	}

	return locations, nil
}
