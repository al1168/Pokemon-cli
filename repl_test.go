package main

import(
	"testing"
	// "fmt"
)
func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: " H1llo wOrld Today is good",
			expected: []string{"h1llo", "world", "today", "is", "good"},
		},
	}
	for caseNum, c := range cases{

		cleanInputResult := cleanInput(c.input)
		for i, word := range cleanInputResult{
			expectedWord := c.expected[i]
			if expectedWord != word{
				t.Errorf("Failed case %v \n failed to match at index %v\n cleaned word: %v expected word: %v", caseNum, i, word, expectedWord)
				return
			}
		}

	}
}
