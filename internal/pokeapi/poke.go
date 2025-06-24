package pokeapi
// package main


import (
	"fmt"
	// "io"
	"net/http"
	"encoding/json"
)

const (
	BASE_URL = "https://pokeapi.co/api/v2/location-area"
)
func GetPokemonLocation(offset int) error{
	fmt.Printf("The offset is : %v\n", offset)
	locationUrl := BASE_URL + "?"+"offset=" + fmt.Sprintf("%d", offset) 
	// var locationJson
	fmt.Printf("The locationUrl is: %v\n", locationUrl)
	resp, err := http.Get(locationUrl)
	if err != nil {
		return fmt.Errorf("fail to make GET request to %v", locationUrl)
	}
	// locationJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("fail to read response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}
	decoder := json.NewDecoder(resp.Body)
	var locationStruct PokemonLocationStruct
	defer resp.Body.Close()
	err = decoder.Decode(&locationStruct)
	if err != nil {
		return fmt.Errorf("fail to unmarshal JSON: %v", err)
	}
	// fmt.Printf("This is :\n %v ", locationStruct)
	// Print the results
	for _, result := range locationStruct.Results {
		fmt.Printf("%v\n", result.Name)
	}
	// fmt.Printf("This is the locationJSON: %s", locationJSON)
	return nil
}