package main

import (
	"time"

	"github.com/havokmoobii/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 1 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: make(map[string]pokemon),
	}
	
	startRepl(cfg)
}
