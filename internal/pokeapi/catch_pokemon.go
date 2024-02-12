package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jmarkIT/pokedexcli/internal/pokecache"
)

func (c *Client) GetPokemon(pokemon string, cache *pokecache.Cache) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + pokemon + "/"

	pokemonResp := RespPokemon{}

	// Look in the cache for the requested data
	if dat, exists := cache.Get(url); exists {
		err := json.Unmarshal(dat, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}
		return pokemonResp, nil
	}

	// If we get here the data isn't in the cache so we need to hit the api
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}

	cache.Add(url, dat)
	return pokemonResp, nil
}
