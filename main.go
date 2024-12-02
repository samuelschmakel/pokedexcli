package main

import (
	"math/rand"
	"time"

	pokeapi "github.com/samuelschmakel/pokedexcli/internal/pokeapi"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	pokeapi.Init()
	startRepl()
}
