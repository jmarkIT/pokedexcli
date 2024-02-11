package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jmarkIT/pokedexcli/internal/pokecache"
)

// Get details of a specific area
func (c *Client) GetArea(cache *pokecache.Cache, areaName string) (RespAreas, error) {
	url := baseURL + "/location-area/" + areaName

	areaResp := RespAreas{}

	// Look in the cache for the requested data
	if dat, exists := cache.Get(url); exists {
		err := json.Unmarshal(dat, &areaResp)
		if err != nil {
			return RespAreas{}, err
		}
		return areaResp, nil
	}

	// If we get here the data isn't in the cache so we need to hit the api
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespAreas{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespAreas{}, err
	}

	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		return RespAreas{}, err
	}

	cache.Add(url, dat)
	return areaResp, nil
}
