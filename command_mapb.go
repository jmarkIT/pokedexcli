package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMapB(configuration *config) error {
	if configuration.previous == nil {
		return errors.New("No previous page locations to display")
	}

	url := configuration.previous

	res, err := http.Get(*url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

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
