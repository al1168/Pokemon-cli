package main

import (
	"fmt"
)

func ExploreCommand(c *config, args ...string) error{
	client := c.pokemonapiClient
	if len(args) == 0 {
		return fmt.Errorf("please Input city name after Explore\nuserInput:%v",nil)
	}
	locationString := args[0]
	pokeInfo, err := client.ListPokemonAtLocation(&locationString)
	if err != nil{
		return fmt.Errorf("failed to retrieve data from ListAtPokemonLocation call, error: %v", err)
	}
	fmt.Printf("Exploring %v\n", locationString)
	for _, pokemonStruct := range pokeInfo.PokemonEncounters{
		pokemonName := pokemonStruct.Pokemon.Name
		fmt.Printf("%v\n", pokemonName)
	}
	return nil
}