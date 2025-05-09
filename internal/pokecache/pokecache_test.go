package pokecache

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	const interval = 5 * time.Second
	type testCase struct {
		name        string
		key         string
		expectedVal []byte
	}

	cases := []testCase{
		{
			name:        "adds new url and returns response",
			key:         "https://pokeapi.co/api/v2/location-areas",
			expectedVal: []byte("pokeapi-response"),
		},
		{
			name:        "adds new url with with empty response and returns it",
			key:         "https://pokeapi.co/api/v2/location-areas",
			expectedVal: []byte{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tc.key, tc.expectedVal)
			actualVal, ok := cache.Get(tc.key)
			if !ok {
				t.Errorf("expected to find a key [%s]: got [%t] want [%t]", tc.key, ok, !ok)
				return
			}
			if string(actualVal) != string(tc.expectedVal) {
				t.Errorf("returned values don't match: got [%v] want [%v]", actualVal, tc.expectedVal)
				return
			}
		})
	}
}

func TestGet(t *testing.T) {
	const interval = 5 * time.Millisecond

	type testCase struct {
		name            string
		key             string
		expectedToExist bool
		expectedValue   []byte
	}

	cases := []testCase{
		{
			name:            "returns existing value",
			key:             "https://pokeapi.co/api/v2/location-areas",
			expectedToExist: true,
			expectedValue:   []byte("pokeapi-response"),
		},
		{
			name:            "returns feedback about key that doesn't exist and empty value",
			key:             "https://pokeapi.co/api/v2/location-areas",
			expectedToExist: false,
			expectedValue:   []byte{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cache := NewCache(interval)
			if tc.expectedToExist {
				cache.Add(tc.key, tc.expectedValue)
			}
			actualVal, ok := cache.Get(tc.key)

			if ok != tc.expectedToExist {
				t.Errorf("existance of a key doesn't match: got [%t] want [%t]", ok, tc.expectedToExist)
				return
			}
			if string(actualVal) != string(tc.expectedValue) {
				t.Errorf("values don't match: got [%v] want %v", actualVal, tc.expectedValue)
				return
			}
		})
	}
}
