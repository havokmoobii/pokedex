package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationResp struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}               `json:"results"`
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.next != "" {
		url = c.next
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var locations LocationResp
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)

	if err != nil {
		return fmt.Errorf("decoder error: %w", err)
	}

	if locations.Next != nil {
		c.next = *locations.Next
	} else {
		c.next = ""
	}
	if locations.Previous != nil {
		c.previous = *locations.Previous
	} else {
		c.previous = ""
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}