package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		cleanIn := cleanInput(scanner.Text())
		if len(cleanIn) == 0 {
			fmt.Println("Missing Command")
			continue
		}
		command, exists := getCommands()[cleanIn[0]]
		if !exists {
			fmt.Println("Unknown Command")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Errorf("error running command: %w", err)
			os.Exit(1)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name	    string
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