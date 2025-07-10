package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
func (c *Client)GetPokemonInfo(pokemonName *string) (PokemonInfo, error) {
	if pokemonName == nil {
		return PokemonInfo{}, fmt.Errorf("error: PokemonName was nil %v", pokemonName)
	}
	client := c.client
	providedPokemonName := *pokemonName
	url := BASE_URL + "pokemon/" + providedPokemonName
	req, err := http.NewRequest("GET",url, nil)
	if err != nil{
		return PokemonInfo{}, fmt.Errorf("error: a problem occured creating request")
	}
	resp, err := client.Do(req)
	if err != nil{
		return PokemonInfo{}, fmt.Errorf("This is an invalid pokemon")
	}
	var pokeinfoContainer PokemonInfo
	data, err := io.ReadAll(resp.Body)
	if err != nil{
		return PokemonInfo{}, fmt.Errorf("have problem reading byte into data, %v", err)
	}
	err = json.Unmarshal(data, &pokeinfoContainer)
	if err != nil{
		return PokemonInfo{}, fmt.Errorf("error from unmsarshaling in getPokemonInfo %v", err)
	}
	return pokeinfoContainer, nil
}