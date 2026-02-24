package pokeapi

type LocationPokemonsResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Pokemons []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
