package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliName string = "Pokedex"

func commandHelp() error {
	fmt.Printf("Welcome to %v! These are the available commands: \n", cliName)
	fmt.Println("help - Show available commands")
	fmt.Println("exit - Closes your connection to ", cliName)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func main() {
	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	commands := map[string]cliCommand{
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
	}

	scanner := bufio.NewScanner(os.Stdin)

	// start REPL loop
	for {
		fmt.Printf("%v > ", cliName)
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if cmdFunc, exists := commands[input]; exists {
			cmdFunc.callback()
			continue
		}

		// Default behavior for unrecognized input
		fmt.Println("Unknown command:", input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input: ", err)
	}
}
