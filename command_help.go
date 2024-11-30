package main

import "fmt"

var cliName string = "Pokedex"

func commandHelp() error {
	fmt.Printf("Welcome to %v! These are the available commands: \n", cliName)
	fmt.Println("")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}
