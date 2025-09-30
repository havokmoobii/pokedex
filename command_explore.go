package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, parameters []string) error {
	if len(parameters) == 0 {
		return errors.New("Missing area name")
	}
	areaName := parameters[0]

	fmt.Printf("Exploring %s...\n", areaName)

	areaResp, err := cfg.pokeapiClient.AreaInformation(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range areaResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}