package main

import (
	"fmt"
	"math/rand"
	"time"
)

func tryCatch(baseExperience int) bool {
	attempt := rand.Float64() * 100
	catchChance := calculateCatchChance(baseExperience)
	return attempt <= catchChance
}

func calculateCatchChance(baseExperience int) float64 {
	const probabilityDivisor = 10.0
	const maxBaseExp = 400.0

	baseExpFloat := float64(baseExperience)
	if baseExperience < 0 {
		return 100.0
	}
	if baseExpFloat > maxBaseExp {
		return 100.0 - maxBaseExp/probabilityDivisor
	}
	return 100.0 - baseExpFloat/probabilityDivisor
}

func createTension() {
	for range 4 {
		fmt.Printf(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
}
