package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
func (c *Client)GetPokemonInfo(pokemonName *string) (PokemonInfo, error) {
	var pokeinfoContainer PokemonInfo
	if pokemonName == nil {
		return PokemonInfo{}, fmt.Errorf("error: PokemonName was nil %v", pokemonName)
	}
	client := c.client
	providedPokemonName := *pokemonName
	cache := c.cache
	catchedData, isInCache, _ := cache.Get(providedPokemonName)
	if isInCache {
		unMarshalerr := json.Unmarshal(catchedData, &pokeinfoContainer)
		if unMarshalerr != nil{
			return PokemonInfo{}, fmt.Errorf("error from unmsarshaling in getPokemonInfo %v", unMarshalerr)
		}
		return pokeinfoContainer, nil
	}
	url := BASE_URL + "pokemon/" + providedPokemonName
	req, err := http.NewRequest("GET",url, nil)
	if err != nil{
		return PokemonInfo{}, fmt.Errorf("error: a problem occured creating request")
	}
	resp, err := client.Do(req)
	if err != nil{
		return PokemonInfo{}, fmt.Errorf("this is an invalid pokemon")
	}
	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil{
		return PokemonInfo{}, fmt.Errorf("have problem reading byte into data, %v", err)
	}
	err = json.Unmarshal(data, &pokeinfoContainer)
	if err != nil{
		return PokemonInfo{}, fmt.Errorf("error from unmsarshaling in getPokemonInfo %v", err)
	}
	cache.Add(providedPokemonName, data)
	return pokeinfoContainer, nil
}