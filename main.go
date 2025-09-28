package main

import (
	"time"

	"github.com/havokmoobii/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	
	startRepl(cfg)
}

// Next time: Try to understand what is being done in the internal files of the solution