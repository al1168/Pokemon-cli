package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

func (c *Client) ListPokemonAtLocation(pageUrl *string) (PokemonAtLocation, error){
	client := c.client
	url := BASE_URL + "location-area/" + *pageUrl
	cache := c.cache
	cachedData, isInCache, _ := cache.Get(url)
	var unmarshaledData PokemonAtLocation
	if isInCache {
		cacheErr := json.Unmarshal(cachedData, &unmarshaledData)
		if cacheErr != nil {
			return PokemonAtLocation{}, fmt.Errorf("error with unmarshaling cache item request: %v", cacheErr)
		}
		return unmarshaledData, nil
	}
	newRequest, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return PokemonAtLocation{}, fmt.Errorf("error occured creating http.NewRequest\n Error: %v", err)
	}
	responsObject, err := client.Do(newRequest)
	if err !=nil{
		return PokemonAtLocation{}, fmt.Errorf("error occured at listPokemonAtLocation making request\n error: %v", err)
	}
	data, err := io.ReadAll(responsObject.Body)
	if err != nil {
		return PokemonAtLocation{}, fmt.Errorf("error occured at listPokemonAtLocation reading res.Body\n error: %v", err)
	}

	err = json.Unmarshal(data, &unmarshaledData)
	if err != nil {
		return PokemonAtLocation{}, fmt.Errorf("error occured at listPokemonAtLocation unmarshaling\n error: %v", err)
	}
	
	return unmarshaledData, nil
}