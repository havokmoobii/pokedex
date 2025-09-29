package main

import "fmt"

func commandHelp(cfg *config, parameters []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	if len(parameters) >= 1 {
		command, exists := getCommands()[parameters[0]]
		if !exists {
			fmt.Println("Unknown Command")
			return nil
		}

		fmt.Printf("%s: %s\n", command.name, command.usage)
		return nil
	}

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}