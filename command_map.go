package main

import (
	"errors"
	"fmt"

	"github.com/jmarkIT/pokedexcli/internal/pokeapi"
	"github.com/jmarkIT/pokedexcli/internal/pokecache"
)

func commandMap(cfg *config, cache *pokecache.Cache, _ string, _ map[string]pokeapi.RespPokemon) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL, cache)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.previousLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, cache *pokecache.Cache, _ string, _ map[string]pokeapi.RespPokemon) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL, cache)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.previousLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
