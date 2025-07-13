package main

import (
	"fmt"
)
func MapCommand(c *config, args...string) error{
	client := c.pokemonapiClient
	var err error 

	responseLoc, err := client.ListLocations(c.nextURL)
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