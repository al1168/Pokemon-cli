package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/al1168/Pokemon-cli/internal/pokecache"
	"io"
)
func (c *Client) ListLocations(pageUrl *string, cache *pokecache.Cache) (PokemonLocationStruct, error){
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
	data, err := io.ReadAll(responsObject.Body)
	if err != nil {
		return PokemonLocationStruct{}, fmt.Errorf("fail to ioRead data, %v", err)
	}
	var locationObj PokemonLocationStruct
	err = json.Unmarshal(data,&locationObj)
	if err != nil{
		return PokemonLocationStruct{}, fmt.Errorf("fail to decode to struct, %v", err)
	}
	cache.Add(url, data)
	return locationObj, nil
}