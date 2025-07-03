package main

import (
	"encoding/json"
	"fmt"
	// "net/http"

	// "github.com/al1168/Pokemon-cli/internal/pokeapi"
	"github.com/al1168/Pokemon-cli/internal/pokeapi"
	"github.com/al1168/Pokemon-cli/internal/pokecache"
)

func MapCommandB(c *config, cache *pokecache.Cache) error{
	client := c.pokemonapiClient
	var err error

	if c.prevUrl == nil || *c.prevUrl == "" {
		fmt.Print("We cannot go back any further\n")
		return nil
	}
	cacheValue, isInCache, _ := cache.Get(*c.prevUrl)
	var responseLoc pokeapi.PokemonLocationStruct
	// if cacheErr != nil{
		// fmt.Printf("%v", cacheErr)
	// }
	if !isInCache {
		responseLoc, err = client.ListLocations(c.prevUrl, cache)
	} else{
		err = json.Unmarshal(cacheValue, &responseLoc)
		// fmt.Printf("FOUND IN %v in CACHE \n", *c.prevUrl)
	}
	if err != nil{
		return fmt.Errorf("failed to get loc structd from ListLocations, error: %v", err)
	}
	
	for _, result := range responseLoc.Results{
		fmt.Printf("%v \n", result.Name)
	}
	fmt.Println("")
	c.nextURL = &responseLoc.Next
	c.prevUrl = &responseLoc.Previous
	return nil
}