package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/jmarkIT/pokedexcli/internal/pokeapi"
	"github.com/jmarkIT/pokedexcli/internal/pokecache"
)

func commandCatch(cfg *config, cache *pokecache.Cache, pokemon string, pokedex map[string]pokeapi.RespPokemon) error {
	if pokemon == "" {
		return errors.New("please provide the name of the pokemon to catch")
	}
	respPokemon, err := cfg.pokeapiClient.GetPokemon(pokemon, cache)
	if err != nil {
		return err
	}

	toCatch := rand.Intn(respPokemon.BaseExperience)
	catchAttempt := rand.Intn(100)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if catchAttempt >= toCatch {
		fmt.Printf("%s was caught!\n", pokemon)
		pokedex[pokemon] = respPokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
