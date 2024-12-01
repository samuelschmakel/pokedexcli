package main

import (
	"os"

	pokeapi "github.com/samuelschmakel/pokedexcli/internal/pokeapi"
)

func commandExit(myConfig *pokeapi.Config) error {
	os.Exit(0)
	return nil
}
