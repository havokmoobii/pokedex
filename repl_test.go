package main

import (
	"testing"
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
		{
			input:	  "",
			expected: []string{},
		},
		{
			input:	  " Help    me   PLZ                              ",
			expected: []string{"help", "me", "plz"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Error: Actual length does not match expected length")
			t.Errorf("Actual: %v   Expected: %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Error: Actual word does not match expected word")
				t.Errorf("Actual: %s   Expected: %s", word, expectedWord)
			}
		}
	}
}

// Work on this a bit more before moving on. More tests to start