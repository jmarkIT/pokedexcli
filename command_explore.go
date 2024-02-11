package main

import (
	"errors"
	"fmt"

	"github.com/jmarkIT/pokedexcli/internal/pokecache"
)

func commandExplore(cfg *config, cache *pokecache.Cache, area string) error {
	if area == "" {
		return errors.New("please provide the area to explore")
	}
	areaResp, err := cfg.pokeapiClient.GetArea(cache, area)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, poke := range areaResp.PokemonEncounters {
		fmt.Printf(" - %s\n", poke.Pokemon.Name)
	}

	return nil
}
