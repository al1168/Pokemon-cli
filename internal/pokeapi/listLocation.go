package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)
func (c *Client) ListLocations(pageUrl *string) (PokemonLocationStruct, error){
	cache := c.cache
	var locationObj PokemonLocationStruct
	url := BASE_URL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	cacheValue, isInCache, _ := cache.Get(url)
	var err error 
	// if in cache return cached item
	if isInCache {
		err = json.Unmarshal(cacheValue, &locationObj)
		if err != nil {
			return PokemonLocationStruct{}, fmt.Errorf("error with unmarshaling cache item request: %v", err)
		}
		return locationObj, nil
	}
	newRequest, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return PokemonLocationStruct{}, fmt.Errorf("error with create request: %v", err)
	}
	responsObject, error := c.client.Do(newRequest)
	if error!= nil{
		return PokemonLocationStruct{}, fmt.Errorf("error Making request, %v", error)
	}
	data, err := io.ReadAll(responsObject.Body)
	if err != nil {
		return PokemonLocationStruct{}, fmt.Errorf("fail to ioRead data, %v", err)
	}
	
	err = json.Unmarshal(data, &locationObj)
	if err != nil{
		return PokemonLocationStruct{}, fmt.Errorf("fail to decode to struct, %v", err)
	}
	cache.Add(url, data)
	return locationObj, nil
}