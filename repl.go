package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	commands := getCommands()

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

		command := cleaned[0]
		fmt.Println(cleaned)

		switch command {
		case "exit":
			commandExit()
		}

		if cmdFunc, exists := commands[command]; exists {
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

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}
