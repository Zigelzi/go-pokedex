package main

import "testing"

func TestCalculateCatchChance(t *testing.T) {
	cases := []struct {
		name                string
		baseExperience      int
		expectedCatchChance float64
	}{
		{
			name:                "returns realistic probability for typical input",
			baseExperience:      120,
			expectedCatchChance: 88,
		},
		{
			name:                "returns max probability for negative input",
			baseExperience:      -100,
			expectedCatchChance: 100,
		},
		{
			name:                "returns max probability for zero input",
			baseExperience:      0,
			expectedCatchChance: 100,
		},
		{
			name:                "returns high probability for small input",
			baseExperience:      1,
			expectedCatchChance: 99.9,
		},
		{
			name:                "returns moderate probability for high input",
			baseExperience:      400,
			expectedCatchChance: 60,
		},
		{
			name:                "returns moderate probability for capped high input",
			baseExperience:      1000,
			expectedCatchChance: 60,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actualCatchChance := calculateCatchChance(tc.baseExperience)
			if actualCatchChance != tc.expectedCatchChance {
				t.Errorf("catch chances don't match: got [%.2f] want [%.2f]",
					actualCatchChance,
					tc.expectedCatchChance)
			}
		})
	}
}
