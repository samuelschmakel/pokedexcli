package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/samuelschmakel/pokedexcli/internal/pokeapi"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	// start REPL loop
	for {
		fmt.Printf("%v > ", cliName)
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		cmdFunc, exists := getCommands()[commandName]
		if exists {
			err := cmdFunc.callback(*&pokeapi.InitialConfig)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			// Default behavior for unrecognized input
			fmt.Println("Unknown command:", input)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input: ", err)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}
