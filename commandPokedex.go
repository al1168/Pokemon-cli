package main

import "fmt"

func PokedexCommand(c *config, args ...string) error{
	if len(args) != 0{
		fmt.Print("This is only to list caught pokemon\n")
		return nil
	}
	fmt.Print("Your Pokedex:")
	for key, _ := range c.pokedex{
		fmt.Printf("  -%v\n", key)
	}
	return nil
}
