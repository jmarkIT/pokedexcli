package main

import (
	"os"

	"github.com/jmarkIT/pokedexcli/internal/pokecache"
)

func commandExit(cfg *config, cache *pokecache.Cache) error {
	os.Exit(0)
	return nil
}
