package main

import (
	"math/rand"
	"testing"
)

func TabuSearchTest(t *testing.T, params ProblemParams) {
	maxIterations := 5000  // you should not need to change this
	limitNotImproved := 10 // you probably should not need to change this
	r := rand.New(rand.NewSource(1))
	evaluate, findNeighborhood, blocks := getSearchParams(params)

	tabuListSize := 20
	initialSolution := GenRandomSolution(params.NumDays, blocks.Weights, r)
	solution := TabuSearch(evaluate, findNeighborhood, initialSolution, maxIterations, tabuListSize, limitNotImproved)

	if solution.Score.Penalty != 0 {
		t.Error("Expected solution penalty to be 0, got ", solution.Score.Penalty)
	}
}

func TestTabuSearch(t *testing.T) {
	params := paramSets[0]
	TabuSearchTest(t, params)
}

func TestTabuSearch2(t *testing.T) {
	params := paramSets[1]
	TabuSearchTest(t, params)
}

func TestTabuSearch3(t *testing.T) {
	params := paramSets[2]
	TabuSearchTest(t, params)
}

func TestTabuSearch4(t *testing.T) {
	params := paramSets[3]
	TabuSearchTest(t, params)
}
