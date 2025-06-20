package main

import (
	"fmt"
	"strings"
)
func cleanInput(text string) []string {

	splittedStrings := strings.Split(text, " ")
	lowerCasedStrings := []string{}
	for _, str := range splittedStrings {
		if str == ""{
			continue
		}
		// lowerCasedStrings[i] = strings.ToLower(str)	
		lowerCasedStrings = append(lowerCasedStrings, strings.ToLower(str))
	}
	fmt.Println(lowerCasedStrings)
	fmt.Println(len(lowerCasedStrings))

	return lowerCasedStrings
}
func main(){
	fmt.Print("Hello, World!")

	myString := "Hello World This Is A tEst"
	cleanInput(myString)
}