package main

import (
	"math/rand"
	"testing"
)

func LAHCTest(t *testing.T, params ProblemParams) {
	evaluate, findNeighborhood, blocks := getSearchParams(params)
	maxIterations := 5000  // you should not need to change this
	limitNotImproved := 10 // you probably should not need to change this
	r := rand.New(rand.NewSource(1))

	listSize := 100
	initialSolution := GenRandomSolution(params.NumDays, blocks.Weights, r)
	solution := LAHC(evaluate, findNeighborhood, initialSolution, maxIterations, listSize, limitNotImproved)

	if solution.Score.Penalty != 0 {
		t.Error("Expected solution penalty to be 0, got ", solution.Score.Penalty)
	}
}

func TestLAHC(t *testing.T) {
	params := paramSets[0]
	LAHCTest(t, params)
}

func TestLAHC2(t *testing.T) {
	params := paramSets[1]
	LAHCTest(t, params)
}

func TestLAHC3(t *testing.T) {
	params := paramSets[2]
	LAHCTest(t, params)
}

func TestLAHC4(t *testing.T) {
	params := paramSets[3]
	LAHCTest(t, params)
}
