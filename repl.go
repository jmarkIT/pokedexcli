package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jmarkIT/pokedexcli/internal/pokeapi"
	"github.com/jmarkIT/pokedexcli/internal/pokecache"
)

func startRepl(cfg *config, cache *pokecache.Cache) {
	reader := bufio.NewScanner(os.Stdin)
	pokedex := make(map[string]pokeapi.RespPokemon)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var subCommand string
		if len(words) > 1 {
			subCommand = words[1]
		} else {
			subCommand = ""
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, cache, subCommand, pokedex)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, cache *pokecache.Cache, area string, pokedex map[string]pokeapi.RespPokemon) error
}

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display's the previous 20 location",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays all the pokemon found in the provided area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the provided pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Get details of the provided pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all of your caught Pokemon",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
