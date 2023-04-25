package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetPokemonByName(
	ctx context.Context,
	pokemonName string,
) (Pokemon, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/api/v2/pokemon/"+pokemonName,
		nil,
	)

	if err != nil {
		return Pokemon{}, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code returned from the pokeapi")
	}

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
