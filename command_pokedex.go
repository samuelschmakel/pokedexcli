package main

import (
	"fmt"

	pokeapi "github.com/samuelschmakel/pokedexcli/internal/pokeapi"
)

func commandPokedex(myConfig *pokeapi.Config, params []string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range pokeapi.CaughtPokemon.Poke {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
