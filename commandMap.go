package main

import (
	"fmt"
	"github.com/al1168/Pokemon-cli/internal/pokecache"
)
func MapCommand(c *config, cache *pokecache.Cache) error{
	client := c.pokemonapiClient
	
	responseLoc, err := client.ListLocations(c.nextURL, cache)
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