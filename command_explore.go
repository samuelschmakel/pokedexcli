package main

import (
	"encoding/json"
	"fmt"

	pokeapi "github.com/samuelschmakel/pokedexcli/internal/pokeapi"
)

func commandExplore(myConfig *pokeapi.Config, params []string) error {
	fmt.Println(len(params))
	if len(params) == 0 {
		return fmt.Errorf("explore requires a parameter")
	}

	// Keep everything below this line
	testURL := "https://pokeapi.co/api/v2/location-area/" + params[0] + "/"

	t, err := pokeapi.GetLocationAreas(testURL)
	if err != nil {
		return err
	}

	var location pokeapi.LocationArea
	err = json.Unmarshal(t, &location)
	if err != nil {
		return err
	}

	for i := 0; i < len(location.PokemonEncounters); i++ {
		fmt.Printf("- %s\n", location.PokemonEncounters[i].Pokemon.Name)
	}
	//fmt.Printf("Here's the JSON: %v\n", location.PokemonEncounters)

	return nil
}
