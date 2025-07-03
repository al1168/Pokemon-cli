package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

func (c Client) listPokemonAtLocation(pageUrl *string) (PokemonAtLocation, error){
	client := c.client
	newRequest, err := http.NewRequest("GET", *pageUrl, nil)
	if err != nil{
		return PokemonAtLocation{}, fmt.Errorf("error occured creating http.NewRequest\n Error: %v", err)
	}
	var unmarshaledData PokemonAtLocation
	responsObject, err := client.Do(newRequest)
	if err !=nil{
		return PokemonAtLocation{}, fmt.Errorf("error occured at listPokemonAtLocation making request\n error: %v", err)
	}
	data, err := io.ReadAll(responsObject.Body)
	if err != nil {
		return PokemonAtLocation{}, fmt.Errorf("error occured at listPokemonAtLocation reading res.Body\n error: %v", err)
	}

	err = json.Unmarshal(data, &unmarshaledData)
	if err != nil {
		return PokemonAtLocation{}, fmt.Errorf("error occured at listPokemonAtLocation unmarshaling\n error: %v", err)
	}
	
	return unmarshaledData, nil
}