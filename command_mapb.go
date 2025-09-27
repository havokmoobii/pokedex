package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapb(c *config) error {
	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(c.previous)
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