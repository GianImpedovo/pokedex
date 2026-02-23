package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(c *config) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	url := c.Previous // Esta es la URL que queremos

	// 1. ¿Está en la cache?
	info, ok := c.pokeapiClient.Get(url)
	var data locationAreaResponse

	if ok {
		// 2. Si está, la usamos
		err := json.Unmarshal(info, &data)
		if err != nil {
			return fmt.Errorf("could not parse response body: %w", err)
		}
	} else {
		// 3. Si no está, vamos a internet
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("could not get location area response: %w", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("could not read response body: %w", err)
		}

		// 4. ¡IMPORTANTÍSIMO! Guardamos en la cache para la próxima vez
		c.pokeapiClient.Add(url, body)

		err = json.Unmarshal(body, &data)
		if err != nil {
			return fmt.Errorf("could not parse response body: %w", err)
		}
	}

	// Actualizamos el estado y mostramos los resultados
	c.Next = data.Next
	c.Previous = data.Previous

	for _, v := range data.Result {
		fmt.Println(v.Name)
	}

	return nil

}
