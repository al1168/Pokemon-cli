package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/al1168/Pokemon-cli/internal/pokeapi"
	"github.com/al1168/Pokemon-cli/internal/pokecache"
	"time"
)

type cliCommand struct{
	name string 
	description string
	callback func(*config, *pokecache.Cache) error
}

type config struct{
	pokemonapiClient pokeapi.Client
	nextURL *string
	prevUrl *string
}

func commandExit(*config, *pokecache.Cache) error{
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config, *pokecache.Cache) error{
	fmt.Print(
`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
return nil
}

func startRepl(){
	cache := pokecache.NewCache(time.Second*10)
	scanner := bufio.NewScanner(os.Stdin)
	// run process to clean out cache
	// go
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
			callback: MapCommand,
		},
		"mapb": {
			name: "mapb",
			description: "Get the prev 10 Pokemon locations",
			callback: MapCommandB,
		},
	}
	myConfig := config{
			pokemonapiClient: pokeapi.NewClient(5 * time.Second),
		}
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
		err := callback(&myConfig, &cache)
		if err != nil{
			fmt.Printf("An error occured %v", err)
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