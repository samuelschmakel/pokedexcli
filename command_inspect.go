package main

import (
	"fmt"

	pokeapi "github.com/samuelschmakel/pokedexcli/internal/pokeapi"
)

func commandInspect(myConfig *pokeapi.Config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("inspect requires a parameter")
	}
	pok, ok := pokeapi.CaughtPokemon.Poke[params[0]]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", params[0])
	fmt.Printf("Height: %d\n", pok.Height)
	fmt.Printf("Weight: %d\n", pok.Weight)
	fmt.Println("Stats:")

	for i := 0; i < len(pok.Stats); i++ {
		fmt.Printf("  -%s: %d\n", pok.Stats[i].Stat.Name, pok.Stats[i].BaseStat)
	}

	println("Types:")

	for i := 0; i < len(pok.Types); i++ {
		fmt.Printf("  - %s\n", pok.Types[i].Type.Name)
	}

	//fmt.Printf("   -hp: ")
	return nil
}
