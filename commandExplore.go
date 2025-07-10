package main

import (
	"fmt"
)

func ExploreCommand(c *config) error{
	client := c.pokemonapiClient
	var locationString *string
	if c.userInput == nil{
		return fmt.Errorf("please Input city name after Explore\nuserInput:%v",nil)
	}
	locationString = c.userInput
	pokeInfo, err := client.ListPokemonAtLocation(locationString)
	if err != nil{
		return fmt.Errorf("failed to retrieve data from ListAtPokemonLocation call, error: %v", err)
	}
	fmt.Printf("Exploring %v\n", *locationString)
	for _, pokemonStruct := range pokeInfo.PokemonEncounters{
		pokemonName := pokemonStruct.Pokemon.Name
		fmt.Printf("%v\n", pokemonName)
	}
	return nil
}