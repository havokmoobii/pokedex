package main

import (
	"fmt"
	"errors"
	"math"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, parameters []string) error {
	if len(parameters) == 0 {
		return errors.New("Missing Pokemon name")
	}
	pokemonName := parameters[0]

	speciesResp, err := cfg.pokeapiClient.SpeciesInformation(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	if !throwBall(speciesResp.CaptureRate) {
		fmt.Println(pokemonName, "escaped!")
		return nil
	}
	fmt.Println(pokemonName, "was caught!")
	cfg.pokedex[pokemonName] = pokemon{pokemonName}
	return nil
}

func throwBall(catchRate int) bool{
	catchFactor := math.Max(math.Floor(float64(catchRate) / 12.75), 1)
	shakeFactor := int(math.Floor(1048560 / math.Floor(math.Sqrt(math.Floor(math.Sqrt(math.Floor(16711680 / catchFactor)))))))

	//Modifier to make it a bit fairer than the Safari Zone
	shakeFactor = int(math.Floor(1.5 * float64(shakeFactor)))

	time.Sleep(time.Second)
	if !shakeCheck(shakeFactor) {
		return false
	}
	fmt.Println("*shake*")

	time.Sleep(time.Second)
	if !shakeCheck(shakeFactor) {
		return false
	}
	fmt.Println("*shake*")

	time.Sleep(time.Second)
	if !shakeCheck(shakeFactor) {
		return false
	}
	fmt.Println("*shake*")

	time.Sleep(time.Second)
	if !shakeCheck(shakeFactor) {
		return false
	}
	
	return true
}

func shakeCheck(shakeFactor int) bool {
	shakeValue := rand.Intn(65535)
	return shakeFactor >= int(shakeValue)
}