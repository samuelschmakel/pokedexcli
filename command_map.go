package main

import (
	"encoding/json"
	"fmt"

	pokeapi "github.com/samuelschmakel/pokedexcli/internal/pokeapi"
)

func commandMap(myConfig *pokeapi.Config) error {
	t, err := pokeapi.GetLocationAreas(myConfig.Next)
	if err != nil {
		fmt.Println("You're at the end!")
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
