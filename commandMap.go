package main

import (
	"fmt"
	"encoding/json"
	"github.com/al1168/Pokemon-cli/internal/pokeapi"
	"github.com/al1168/Pokemon-cli/internal/pokecache"
)
func MapCommand(c *config, cache *pokecache.Cache) error{
	client := c.pokemonapiClient
	var err error 
	var nextUrl string
	if c.nextURL != nil{
		nextUrl = *c.nextURL
	}
	cacheValue, isInCache, _ := cache.Get(nextUrl)
	var responseLoc pokeapi.PokemonLocationStruct
	if isInCache {
		// fmt.Printf("%v found in Cache\n", nextUrl)
		json.Unmarshal(cacheValue, &responseLoc)
	} else {
		// fmt.Printf("Cache Error: %v\n", cacheErr)
		responseLoc, err = client.ListLocations(c.nextURL, cache)
	}
	if err != nil{
		return fmt.Errorf("failed to get loc structd from ListLocations, error: %v", err)
	}
	
	for _, result := range responseLoc.Results{
		fmt.Printf("%v \n", result.Name)
	}
	c.nextURL = &responseLoc.Next
	c.prevUrl = &responseLoc.Previous
	return nil
}