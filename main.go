package main

import (
	"fmt"
	"github.com/goml/gobrain"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Seed random
	rand.Seed(time.Now().Unix())

	// Create new AI network
	brain := &gobrain.FeedForward{}

	// Feed brain with some math data
	const dataAmount = 1000 // Amount of examples
	const maxNumber = 10
	const divider = float64(maxNumber + maxNumber)
	mathExamples := make([][][]float64, dataAmount)

	for x := 0; x < dataAmount; x++ {
		a := rand.Intn(maxNumber)
		b := rand.Intn(maxNumber)

		// Add sum example
		mathExamples[x] = [][]float64{{float64(a) / divider, float64(b) / divider}, {float64(a+b) / divider}}
	}

	brain.Init(2, 2, 1)
	brain.Train(mathExamples, dataAmount, 0.6, 0.4, true)
	testAdding(brain, 2, 2, divider)
	testAdding(brain, 5, 5, divider)
	testAdding(brain, 8, 2, divider)
	testAdding(brain, 10, 0, divider)
	testAdding(brain, 3, 2, divider)
	testAdding(brain, 0, 7, divider)
	testAdding(brain, 1, 2, divider)
}

func testAdding(brain *gobrain.FeedForward, a, b int, divider float64) {
	result := brain.Update([]float64{float64(a) / divider, float64(b) / divider})
	resultInterpolated := result[0] * divider
	fmt.Printf("%d + %d = %d (%f)\n", a, b, int(math.Round(resultInterpolated)), resultInterpolated)
}
