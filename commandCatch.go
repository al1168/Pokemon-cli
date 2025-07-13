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

func CatchPokemonCommand(c *config, args ...string) error{
	
	if len(args) == 0 {
		return fmt.Errorf("please Input city name after Explore\nuserInput:%v",nil)
	}
	pokemonName := args[0]
	pokemonInfo, err := c.pokemonapiClient.GetPokemonInfo(&pokemonName)
	if err != nil{
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	basePokemonExperience := pokemonInfo.BaseExperience
	chance := 1 - (float64(basePokemonExperience) /float64(UPPERBOUND))
	isCaught := generateCatchResult(chance)
	
	if isCaught {
		fmt.Printf("%v was caught!\n",pokemonName)
		c.pokedex[pokemonName] = true
		fmt.Printf("Pokedex:%v\n", c.pokedex)
	} else {
		fmt.Printf("%v escaped!\n",pokemonName)
	}
	return nil
}