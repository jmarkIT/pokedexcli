package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(configuration *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if configuration.next != "" {
		url = configuration.next
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	var areas PokeDexArea
	if err := json.Unmarshal(body, &areas); err != nil {
		return err
	}

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	configuration.next = areas.Next
	configuration.previous = areas.Previous
	return nil
}

type PokeDexArea struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
