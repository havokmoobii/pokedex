package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type config struct {
	next     string
	previous string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	c := config{"", ""}
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
		err := command.callback(&c)
		if err != nil {
			fmt.Println("error running command:", err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name	    string
	description string
	callback    func(*config) error
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
			description: "Display the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}