package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

//SpeciesInformation -
func (c *Client) SpeciesInformation(pokemonName string) (ResponseShallowPokemonSpecies, error) {
	url := baseURL + "/pokemon-species/" + pokemonName
	
	cachedResp, exists := c.pokeapiCache.Get(url)

	if exists {
		speciesResp := ResponseShallowPokemonSpecies{}
		err := json.Unmarshal(cachedResp, &speciesResp)
		if err != nil {
			return ResponseShallowPokemonSpecies{}, err
		}
		return speciesResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseShallowPokemonSpecies{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseShallowPokemonSpecies{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseShallowPokemonSpecies{}, err
	}

	c.pokeapiCache.Add(url, dat)

	speciesResp := ResponseShallowPokemonSpecies{}
	err = json.Unmarshal(dat, &speciesResp)
	if err != nil {
		return ResponseShallowPokemonSpecies{}, err
	}
	
	return speciesResp, nil
}

//PokemonInformation -
func (c *Client) PokemonInformation(pokemonName string) (ResponseShallowPokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemonName
	
	cachedResp, exists := c.pokeapiCache.Get(url)

	if exists {
		pokemonResp := ResponseShallowPokemonInfo{}
		err := json.Unmarshal(cachedResp, &pokemonResp)
		if err != nil {
			return ResponseShallowPokemonInfo{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseShallowPokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseShallowPokemonInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseShallowPokemonInfo{}, err
	}

	c.pokeapiCache.Add(url, dat)

	pokemonResp := ResponseShallowPokemonInfo{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return ResponseShallowPokemonInfo{}, err
	}
	
	return pokemonResp, nil
}