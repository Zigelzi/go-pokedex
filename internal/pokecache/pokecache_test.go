package pokecache

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	const testLifetime = 5 * time.Second
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
			cache := NewCache(testLifetime)
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
	const testLifetime = 5 * time.Millisecond

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
			cache := NewCache(testLifetime)
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

func TestReadLoop(t *testing.T) {
	const testLifetime = 5 * time.Millisecond

	type testCase struct {
		name            string
		key             string
		waitDuration    time.Duration
		expectedToExist bool
		expectedValue   []byte
	}

	cases := []testCase{
		{
			name:            "returns cache hit for key with 0 ms age",
			key:             "https://pokeapi.co/api/v2/location-areas",
			waitDuration:    0,
			expectedToExist: true,
			expectedValue:   []byte("pokeapi-response"),
		},
		{
			name:            "returns cache miss for key with 10 ms age",
			key:             "https://pokeapi.co/api/v2/location-areas",
			waitDuration:    10,
			expectedToExist: false,
			expectedValue:   []byte{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cache := NewCache(testLifetime)
			cache.Add(tc.key, tc.expectedValue)
			time.Sleep(tc.waitDuration * time.Millisecond)
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
