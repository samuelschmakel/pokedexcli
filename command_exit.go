package main

import (
	"fmt"
	"os"

	pokeapi "github.com/samuelschmakel/pokedexcli/internal/pokeapi"
)

func commandExit(myConfig *pokeapi.Config, params []string) error {
	if len(params) > 0 {
		return fmt.Errorf("No additional parameters permitted for exit")
	}
	os.Exit(0)
	return nil
}
