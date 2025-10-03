package main

import(
	"fmt"
	"errors"
)


func commandInspect(cfg *config, parameters []string) error {
	if len(parameters) == 0 {
		return errors.New("Missing Pokemon name")
	}
	pokemonName := parameters[0]

	_, exists := cfg.pokedex[pokemonName]
	if !exists {
		fmt.Println(pokemonName, "is not registered in your Pokedex!")
		fmt.Println("Use catch", pokemonName, "to try to catch it!")
		return nil
	}

	fmt.Println("Name:", pokemonName)
	fmt.Println("Height:", cfg.pokedex[pokemonName].height)
	fmt.Println("Weight:", cfg.pokedex[pokemonName].weight)
	fmt.Println("Stats:")

	for _, statInfo := range cfg.pokedex[pokemonName].stats {
		fmt.Printf("  -%s: %v\n", statInfo.name, statInfo.value)
	}
	fmt.Println("Types:")
	for _, pokemonType := range cfg.pokedex[pokemonName].types {
		fmt.Println("  -", pokemonType)
	}
	fmt.Println(cfg.pokedex[pokemonName].flavor)
	return nil
}

