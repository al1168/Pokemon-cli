package main

import (
"fmt"
)


func InspectCommand(c *config, args ...string) error{
	if len(args) == 0{
		return fmt.Errorf("no pokemon name to inspect")
	}
	
	pokemonName := args[0]
	_, exists := c.pokedex[pokemonName]
	
	if !exists{
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	pokemonInfo, error := c.pokemonapiClient.GetPokemonInfo(&pokemonName)
	if error != nil {
		return error
	}

	fmt.Printf("Name: %v\n\nHeight: %v\nWeight: %v", pokemonInfo.Name, pokemonInfo.Height,pokemonInfo.Weight)
	fmt.Print("Stats:\n")
	for _, value := range pokemonInfo.Stats{
		fmt.Printf("  -%v: %v\n", value.Stat.Name, value.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, value := range pokemonInfo.Types{
		fmt.Printf("  -%v\n", value.Type.Name)
	}
	return nil
}