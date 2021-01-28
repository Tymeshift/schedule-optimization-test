package main

import (
	"math/rand"
	"testing"
)

func BenchmarkTabuSearch(b *testing.B) {
	params := paramSets[3]
	evaluate, findNeighborhood, blocks := getSearchParams(params)
	maxIterations := 5000  // you should not need to change this
	limitNotImproved := 10 // you probably should not need to change this
	r := rand.New(rand.NewSource(1))

	for i := 0; i < b.N; i++ {
		tabuListSize := 5
		initialSolution := GenRandomSolution(params.NumDays, blocks.Weights, r)
		solution := TabuSearch(evaluate, findNeighborhood, initialSolution, maxIterations, limitNotImproved, tabuListSize)
		if solution.Score.Penalty != 0 {
			b.Error("Expected solution penalty to be 0, got ", solution.Score.Penalty)
		}
	}
}

func BenchmarkLAHC(b *testing.B) {
	params := paramSets[3]
	evaluate, findNeighborhood, blocks := getSearchParams(params)
	maxIterations := 5000  // you should not need to change this
	limitNotImproved := 10 // you probably should not need to change this
	r := rand.New(rand.NewSource(1))

	for i := 0; i < b.N; i++ {
		listSize := 10
		initialSolution := GenRandomSolution(params.NumDays, blocks.Weights, r)
		solution := LAHC(evaluate, findNeighborhood, initialSolution, maxIterations, listSize, limitNotImproved)
		if solution.Score.Penalty != 0 {
			b.Error("Expected solution penalty to be 0, got ", solution.Score.Penalty)
		}
	}
}

func BenchmarkSimAnneal(b *testing.B) {
	params := paramSets[3]
	evaluate, findNeighborhood, blocks := getSearchParams(params)
	limitNotImproved := 10 // you probably should not need to change this
	r := rand.New(rand.NewSource(1))

	for i := 0; i < b.N; i++ {
		temp := 99999.0
		minTemp := 0.0001
		alpha := 0.99
		strategy := "percentage" // linear, slow, percentage
		initialSolution := GenRandomSolution(params.NumDays, blocks.Weights, r)
		solution := SimAnneal(evaluate, findNeighborhood, initialSolution, limitNotImproved, temp, minTemp, alpha, strategy, r)
		if solution.Score.Penalty != 0 {
			b.Error("Expected solution penalty to be 0, got ", solution.Score.Penalty)
		}
	}
}

// RandomWalk is expected to fail in most cases
func BenchmarkRandom(b *testing.B) {
	params := paramSets[3]
	evaluate, _, blocks := getSearchParams(params)
	maxIterations := 5000  // you should not need to change this
	limitNotImproved := 10 // you probably should not need to change this
	r := rand.New(rand.NewSource(1))

	for i := 0; i < b.N; i++ {
		initialSolution := GenRandomSolution(params.NumDays, blocks.Weights, r)
		solution := Random(evaluate, initialSolution, params, blocks, maxIterations, limitNotImproved, r)
		if solution.Score.Penalty != 0 {
			b.Error("Expected solution penalty to be 0, got ", solution.Score.Penalty)
		}
	}
}
