package main

import (
	"encoding/json"
	"fmt"

	"math/rand"

	pokeapi "github.com/samuelschmakel/pokedexcli/internal/pokeapi"
)

func commandCatch(myConfig *pokeapi.Config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("catch requires a parameter")
	}
	myURL := "https://pokeapi.co/api/v2/pokemon/" + params[0] + "/"

	t, err := pokeapi.GetPokemonInfo(myURL)
	if err != nil {
		return err
	}

	var pokemon pokeapi.Pokemon
	err = json.Unmarshal(t, &pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", params[0])

	roll := rand.Intn(101)
	defaultCatchRate := 85
	if pokemon.BaseExperience >= 300 {
		defaultCatchRate = 15
	} else if pokemon.BaseExperience >= 200 {
		defaultCatchRate = 40
	} else if pokemon.BaseExperience >= 100 {
		defaultCatchRate = 60
	} else if pokemon.BaseExperience > 50 && pokemon.BaseExperience < 100 {
		defaultCatchRate = 75
	}

	if roll <= defaultCatchRate {
		// The pokemon was caught
		fmt.Printf("%s was caught!\n", params[0])

		// Add the caught pokemon to the map
		// I should lock and unlock this!
		pokeapi.CaughtPokemon.Poke[params[0]] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", params[0])
	}

	return nil

}
