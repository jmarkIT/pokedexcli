package main

import (
	"errors"
	"fmt"

	"github.com/jmarkIT/pokedexcli/internal/pokeapi"
	"github.com/jmarkIT/pokedexcli/internal/pokecache"
)

func commandInspect(cfg *config, cache *pokecache.Cache, pokemon string, pokedex map[string]pokeapi.RespPokemon) error {
	if pokemon == "" {
		return errors.New("please provide a pokemon to inspect")
	}
	if _, ok := pokedex[pokemon]; !ok {
		err := fmt.Sprintf("you haven't caught a %s yet!\ntry `catch %s`", pokemon, pokemon)
		return errors.New(err)
	}

	pdat := pokedex[pokemon]

	fmt.Printf("Name: %s\n", pokemon)
	fmt.Printf("Height: %d\n", pdat.Height)
	fmt.Printf("Weight: %d\n", pdat.Weight)
	fmt.Println("Stats:")
	for _, stat := range pdat.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pdat.Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}

	return nil
}
