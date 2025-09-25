package main

import (
	"testing"
	"fmt"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
	input    string
	expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:	  "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Something bad happened")
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			fmt.Println(actual[1])
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("A word was wrong")
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}

// Work on this a bit more before moving on. More tests to start