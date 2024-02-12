package main

import (
	"errors"
	"fmt"

	"github.com/jmarkIT/pokedexcli/internal/pokeapi"
	"github.com/jmarkIT/pokedexcli/internal/pokecache"
)

func commandPokedex(cfg *config, cache *pokecache.Cache, _ string, dex map[string]pokeapi.RespPokemon) error {
	if len(dex) == 0 {
		return errors.New("you haven't caught any pokemon yet")
	}
	fmt.Println("Your Pokedex:")
	for _, p := range dex {
		fmt.Printf("  - %s\n", p.Name)
	}

	return nil
}
