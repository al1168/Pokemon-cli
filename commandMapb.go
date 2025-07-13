package main

import (
	"fmt"
)

func MapCommandB(c *config, arg ...string) error{
	client := c.pokemonapiClient
	var err error

	if c.prevUrl == nil || *c.prevUrl == "" {
		fmt.Print("We cannot go back any further\n")
		return nil
	}
	responseLoc, err := client.ListLocations(c.prevUrl)
	
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