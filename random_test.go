package main

import (
	"math/rand"
	"testing"
)

// Random is expected to fail in most cases

func RandomTest(t *testing.T, params ProblemParams) {
	evaluate, _, blocks := getSearchParams(params)
	maxIterations := 5000  // you should not need to change this
	limitNotImproved := 10 // you probably should not need to change this
	r := rand.New(rand.NewSource(1))

	initialSolution := GenRandomSolution(params.NumDays, blocks.Weights, r)
	solution := Random(evaluate, initialSolution, params, blocks, maxIterations, limitNotImproved, r)

	if solution.Score.Penalty != 0 {
		t.Error("Expected solution penalty to be 0, got ", solution.Score.Penalty)
	}
}

func TestRandom(t *testing.T) {
	params := paramSets[0]
	RandomTest(t, params)
}

func TestRandom1(t *testing.T) {
	params := paramSets[1]
	RandomTest(t, params)
}

func TestRandom2(t *testing.T) {
	params := paramSets[2]
	RandomTest(t, params)
}

func TestRandom3(t *testing.T) {
	params := paramSets[3]
	RandomTest(t, params)
}
