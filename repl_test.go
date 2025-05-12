package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	type testCase struct {
		name          string
		input         string
		expectedWords []string
	}

	testCases := []testCase{
		{
			name:          "splits words separated by space",
			input:         "caterpie blastoise wartortle",
			expectedWords: []string{"caterpie", "blastoise", "wartortle"},
		},
		{
			name:          "trims leading and trailing whitespaces",
			input:         "   hello world      ",
			expectedWords: []string{"hello", "world"},
		},
		{
			name:          "lowercases all words",
			input:         "BULBASAUR CHARIZARD metaphod",
			expectedWords: []string{"bulbasaur", "charizard", "metaphod"},
		},
		{
			name:          "returns empty slice when input is empty",
			input:         "",
			expectedWords: []string{""},
		},
		{
			name:          "returns empty slice when input is only spaces",
			input:         "               ",
			expectedWords: []string{""},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualWords := cleanInput(tc.input)
			if len(actualWords) != len(tc.expectedWords) {
				t.Fatalf("length of returned words don't match: got [%d] want [%d]", len(actualWords), len(tc.expectedWords))
			}

			for i := range actualWords {
				if actualWords[i] != tc.expectedWords[i] {
					t.Errorf("words don't match: got [%s] want [%s]", actualWords[i], tc.expectedWords[i])
				}
			}
		})
	}
}

func TestGetCommandArgument(t *testing.T) {
	type testCase struct {
		name             string
		input            []string
		expectedArgument string
	}

	testCases := []testCase{
		{
			name:             "returns argument when input contains 2 words",
			input:            []string{"explore", "mt-coronet"},
			expectedArgument: "mt-coronet",
		},
		{
			name:             "returns argument when input contains over 2 words",
			input:            []string{"explore", "mt-coronet", "another", "word"},
			expectedArgument: "mt-coronet",
		},
		{
			name:             "returns empty string when input contains 1 word",
			input:            []string{"explore"},
			expectedArgument: "",
		},
		{
			name:             "returns empty string when input contains no words",
			input:            []string{},
			expectedArgument: "",
		},
		{
			name:             "returns empty string when input is nil",
			input:            nil,
			expectedArgument: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			argument := getCommandArgument(tc.input)
			if argument != tc.expectedArgument {
				t.Errorf("arguments don't match: got [%s] want [%s]", argument, tc.expectedArgument)
				return
			}

		})
	}
}
