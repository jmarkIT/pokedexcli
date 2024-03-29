package main

import (
	"fmt"

	"github.com/jmarkIT/pokedexcli/internal/pokeapi"
	"github.com/jmarkIT/pokedexcli/internal/pokecache"
)

func commandHelp(cfg *config, cache *pokecache.Cache, _ string, _ map[string]pokeapi.RespPokemon) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
