package main

import (
	"fmt"
	"math/rand"
)
const UPPERBOUND = 650
func generateCatchResult(pZero float64) bool{
	randomFloat := rand.Float64()
	if randomFloat < pZero{
		return true // caught
	}
	return false // didn't catch
}

func CatchPokemonCommand(c *config) error{
	if c.userInput == nil{
		return fmt.Errorf("please Input city name after Explore\nuserInput:%v",nil)
	}
	if c.userInput == nil {
		return fmt.Errorf("error occured, userInput is nil in CatchPokemon")
	}
	pokemonName := *c.userInput
	pokemonInfo, err := c.pokemonapiClient.GetPokemonInfo(c.userInput)
	if err != nil{
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	basePokemonExperience := pokemonInfo.BaseExperience
	chance := 1 - (float64(basePokemonExperience) /float64(UPPERBOUND))
	isCaught := generateCatchResult(chance)
	// fmt.Printf("This is the chance: %v base exp is: %v\n ", chance, basePokemonExperience)
	if isCaught {
		fmt.Printf("%v was caught!\n",pokemonName)
		c.pokedex[pokemonName] = true
		fmt.Printf("Pokedex:%v\n", c.pokedex)
	} else {
		fmt.Printf("%v escaped!\n",pokemonName)
	}
	return nil
}