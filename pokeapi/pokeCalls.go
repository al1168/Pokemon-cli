package pokeapi

import (
	// "encoding/json"
	"fmt"
	"io"
	"net/http"
)

type config struct{
	offset int
}
const (
	BASE_URL = "https://pokeapi.co/api/v2/location-area"
)
func getPokemonLocation( c *config ) error{
	offset := c.offset
	fmt.Printf("The offset is : %v", offset)
	locationUrl := BASE_URL + "?"+"offset=" + fmt.Sprintf("%d", offset) 
	// var locationJson
	resp, err := http.Get(locationUrl)
	if err != nil {
		return fmt.Errorf("Fail to make GET request to %v", locationUrl)
	}
	locationJSON, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("This is locationJSON:\n %v ", locationJSON)
	return nil
}