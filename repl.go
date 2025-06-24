package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/al1168/Pokemon-cli/internal/pokeapi"
)

type cliCommand struct{
	name string 
	description string
	callback func(*config) error
}

type config struct{
	offset int
}

func commandExit(*config) error{
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error{
	fmt.Print(
`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
return nil
}
func commandMap(c *config )error{
	pokeapi.GetPokemonLocation(c.offset)
	return nil
}
func startRepl(){
	scanner := bufio.NewScanner(os.Stdin)
	cliCommandMap:= map[string]cliCommand{
		"exit" : {
			name: "exit",
			description:"Closes the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Print help menu",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Get the next 10 Pokemon locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Get the prev 10 Pokemon locations",
			callback: commandMap,
		},
	}
	startingOffset := 0
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputtedString := scanner.Text()
		cleanedInput := cleanInput(inputtedString)
		if len(cleanedInput) == 0{
			continue
		}
		firstWord := cleanedInput[0]
		fmt.Printf("First word is: %v\n", firstWord)
		if _, ok := cliCommandMap[firstWord]; !ok{
			fmt.Printf("Command %v not found in Map", firstWord)
			continue
		}
		
		command := cliCommandMap[firstWord]
		callback := command.callback
		myConfig := config{
			offset: startingOffset,
		}

		if command.name == "mapb"{
			if startingOffset == 0{
				fmt.Print("you're on the first page")
				continue
			}
			startingOffset = max(0, startingOffset - 40)
			fmt.Printf("%v",startingOffset)
			myConfig.offset = startingOffset
		}
		callback(&myConfig)
		if command.name == "map"{
			startingOffset += 20
		}

	}
}
func cleanInput(text string) []string {

	splittedStrings := strings.Split(text, " ")
	lowerCasedStrings := []string{}
	for _, str := range splittedStrings {
		if str == ""{
			continue
		}
		lowerCasedStrings = append(lowerCasedStrings, strings.ToLower(str))
	}
	return lowerCasedStrings
}