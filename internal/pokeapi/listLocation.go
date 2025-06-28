package pokeapi
import (
	"fmt"
	"net/http"
	"encoding/json"
)
func (c *Client) ListLocations(pageUrl *string) (PokemonLocationStruct, error){
	// nextUrl := c.nextURL
	url := BASE_URL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	newRequest, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return PokemonLocationStruct{}, fmt.Errorf("error with create request: %v", err)
	}
	responsObject, error := c.client.Do(newRequest)
	if error!= nil{
		return PokemonLocationStruct{}, fmt.Errorf("error Making request, %v", error)
	}
	decoder := json.NewDecoder(responsObject.Body)
	var locationObj PokemonLocationStruct
	err = decoder.Decode(&locationObj)
	if err != nil{
		return PokemonLocationStruct{}, fmt.Errorf("fail to decode to struct, %v", err)
	}

	return locationObj, nil
}