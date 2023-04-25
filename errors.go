package client

import (
	"errors"
	"fmt"
)

var (
	ErrorFetchingPokemon = errors.New("Failed to fetch pokemon")
)

type PokemonFetchErr struct {
	StatusCode 	int
	Message 	string
}

func(e PokemonFetchErr) Error() string {
	return fmt.Sprintf("failed to fetch pokemon: %s with status code %d", e.Message, e.StatusCode)
}