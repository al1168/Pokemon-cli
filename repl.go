package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"Pokemon-cli/pokeapi"
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
	pokeCalls.getPokemonLocation(c)
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
		if _, ok := cliCommandMap[firstWord]; !ok{
			fmt.Printf("Command %v not found in Map", firstWord)
			continue
		}
		command := cliCommandMap[firstWord]
		callback := command.callback
		myConfig := config{
			offset: startingOffset,
		}
		if command.name == "bmap"{
			startingOffset = max(0, startingOffset - 10)
		}
		callback(&myConfig)
		if command.name == "map"{
			startingOffset += 10
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