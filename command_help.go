package main

import "fmt"

var cliName string = "Pokedex"

func commandHelp() error {
	fmt.Printf("Welcome to %v! These are the available commands: \n", cliName)
	fmt.Println("help - Show available commands")
	fmt.Println("exit - Closes your connection to ", cliName)
	return nil
}
