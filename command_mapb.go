package main

import (
	"encoding/json"
	"fmt"

	pokeapi "github.com/samuelschmakel/pokedexcli/internal/pokeapi"
)

func commandMapb(myConfig *pokeapi.Config, params []string) error {
	if len(params) > 0 {
		return fmt.Errorf("No parameters supported for mapb")
	}
	t, err := pokeapi.GetLocationAreas(myConfig.Previous)
	if err != nil {
		fmt.Println("You're at the beginning!")
		return err
	}

	//locationArea := pokeapi.LocationAreas{}
	var areas pokeapi.LocationAreas
	err = json.Unmarshal(t, &areas)
	if err != nil {
		return err
	}
	for i := 0; i < len(areas.Results); i++ {
		result := areas.Results[i]
		fmt.Println(result.Name)
	}
	myConfig.Next = areas.Next
	str, ok := areas.Previous.(string)
	if !ok {
		myConfig.Previous = ""
	} else {
		myConfig.Previous = str
	}

	return nil
}
