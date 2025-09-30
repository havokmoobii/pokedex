package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"

	"github.com/havokmoobii/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command, exists := getCommands()[words[0]]
		if !exists {
			fmt.Println("Unknown Command")
			continue
		}

		parameters := make([]string, 0)
		if len(words) > 1 {
			parameters = words[1:len(words)]
		}

		err := command.callback(cfg, parameters)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name	    string
	description string
	usage       string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message. Enter help <command> for additional information",
			usage:       `help - Lists all commands and gives basic usage info
      help <command> - Gives additional usage info on specific command if applicable`,
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			usage:       "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 location areas in the Pokemon world",
			usage:       `map - Display the next 20 location areas in the Pokemon world
     map <number> <offset>
         <number> - (NYI)Adjusts how many locations are shown at once
         <offset> - (NYI)Adjusts the starting point in the list of locations`,
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas in the Pokemon world",
			usage:       `map - Display the previous 20 location areas in the Pokemon world
     mapb <number> <offset>
          <number> - (NYI)Adjusts how many locations are shown at once
          <offset> - (NYI)Adjusts the starting point in the list of locations`,
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display all pokemon that can be found in a given area",
			usage:       `explore <location> - Displays all Pokemon that can be found at <location>
                 <location> - Map name that can be found with the map and mapb commands`,
			callback:    commandExplore,
		},
	}
}