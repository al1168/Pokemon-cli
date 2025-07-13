package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/al1168/Pokemon-cli/internal/pokeapi"
	"time"
)

type cliCommand struct{
	name string 
	description string
	callback func(*config, ...string) error
}

type config struct{
	pokemonapiClient pokeapi.Client
	nextURL *string
	prevUrl *string
	pokedex map[string]bool
}

func commandExit(c *config, args ...string) error{
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args ...string) error{
	fmt.Print(
`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
return nil
}

func startRepl(){
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
		"explore": {
			name: "explore",
			description: "Get all pokemon in an area",
			callback: ExploreCommand,
		},
		"catch":{
			name: "catch",
			description:"catch specific pokemon",
			callback: CatchPokemonCommand,
		},
		"inspect": {
			name: "inspect",
			description: "inspect pokemon stats and types",
			callback: InspectCommand,
		},
		"pokedex": {
			name: "pokedex",
			description: "list all pokemon caught",
			callback: PokedexCommand,
		},
	}
	myConfig := config{
			pokemonapiClient: pokeapi.NewClient(5 * time.Second, 10 *time.Second),
			pokedex: make(map[string]bool),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputtedString := scanner.Text()
		cleanedInput := cleanInput(inputtedString)
		if len(cleanedInput) == 0{
			continue
		}
		commandWord := cleanedInput[0]
		
		args := []string{}
		if len(cleanedInput) >= 2{
			args = cleanedInput[1:]
		}
		fmt.Printf("First word is: %v\n", commandWord)
		if _, ok := cliCommandMap[commandWord]; !ok{
			fmt.Printf("Command %v not found in Map\n", commandWord)
			continue
		}
		
		command := cliCommandMap[commandWord]
		callback := command.callback
		err := callback(&myConfig, args...)
		if err != nil{
			fmt.Printf("An error occured %v\n", err)
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