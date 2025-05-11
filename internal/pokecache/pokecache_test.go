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

			if actualVal == nil {
				t.Errorf("expected cache to contain the value for key [%s]: got [nil] want [%v]", tc.key, tc.expectedVal)
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
			name:            "returns cache hit for existing key",
			key:             "https://pokeapi.co/api/v2/location-areas",
			expectedToExist: true,
			expectedValue:   []byte("pokeapi-response"),
		},
		{
			name:            "returns cache miss for non-existing key",
			key:             "https://pokeapi.co/api/v2/location-areas",
			expectedToExist: false,
			expectedValue:   nil,
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
				t.Errorf("existence of a key doesn't match: got [%t] want [%t]", ok, tc.expectedToExist)
				return
			}
			if tc.expectedToExist {
				if string(actualVal) != string(tc.expectedValue) {
					t.Errorf("values don't match: got [%v] want %v", actualVal, tc.expectedValue)
					return
				}
			} else {
				if actualVal != nil {
					t.Errorf("values don't match for cache miss: got [%v] want [nil]", actualVal)
				}
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
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
			name:            "returns cache hit for key which age is less than cache lifetime",
			key:             "https://pokeapi.co/api/v2/location-areas",
			waitDuration:    0,
			expectedToExist: true,
			expectedValue:   []byte("pokeapi-response"),
		},
		{
			name:            "returns cache hit for key which age is equal to cache lifetime",
			key:             "https://pokeapi.co/api/v2/location-areas",
			waitDuration:    5,
			expectedToExist: true,
			expectedValue:   []byte("pokeapi-response"),
		},
		{
			name:            "returns cache miss for key which age is greated than cache lifetime",
			key:             "https://pokeapi.co/api/v2/location-areas",
			waitDuration:    10,
			expectedToExist: false,
			expectedValue:   nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cache := NewCache(testLifetime)
			cache.Add(tc.key, tc.expectedValue)
			time.Sleep(tc.waitDuration * time.Millisecond)
			actualVal, ok := cache.Get(tc.key)

			if ok != tc.expectedToExist {
				t.Errorf("existence of a key doesn't match: got [%t] want [%t]", ok, tc.expectedToExist)
				return
			}
			if tc.expectedToExist {
				if string(actualVal) != string(tc.expectedValue) {
					t.Errorf("values don't match: got [%v] want %v", actualVal, tc.expectedValue)
					return
				}
			} else {
				if actualVal != nil {
					t.Errorf("values don't match for cache miss: got [%v] want [nil]", actualVal)
				}
			}
		})
	}
}
