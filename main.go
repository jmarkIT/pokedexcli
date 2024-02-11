package main

import (
	"time"

	"github.com/jmarkIT/pokedexcli/internal/pokeapi"
	"github.com/jmarkIT/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{pokeapiClient: pokeClient}
	cache := pokecache.NewCache(5 * time.Minute)
	startRepl(cfg, &cache)
}
