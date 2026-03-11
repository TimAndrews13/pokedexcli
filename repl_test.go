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
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "  THIS IS    A TEST    CASE   FOR MY GO    FUNCTION",
			expected: []string{"this", "is", "a", "test", "case", "for", "my", "go", "function"},
		},
		{
			input:    " bLaStOiSe vEnUsAuR cHaRiZaRd         mEw  mEwTwO",
			expected: []string{"blastoise", "venusaur", "charizard", "mew", "mewtwo"},
		},
		{
			input:    "Ash Misty Brock Jessie James",
			expected: []string{"ash", "misty", "brock", "jessie", "james"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("Length of Expected Output does not equal Length of cleanInput Output\nExpected Output: %v\ncleanInput Output: %v", c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected Word is not the same as Word Output by cleanInput\nExpected Word: %v\ncleanInput Word Output: %v", &expectedWord, word)
			}
		}
	}
}
