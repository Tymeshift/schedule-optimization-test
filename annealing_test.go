package main

import (
	"math/rand"
	"testing"
)

func SimAnnealTest(t *testing.T, params ProblemParams) {
	evaluate, findNeighborhood, blocks := getSearchParams(params)

	r := rand.New(rand.NewSource(1))

	limitNotImproved := 100

	temp := 99999.0
	minTemp := 0.0001
	alpha := 0.9
	strategy := "slow" // linear, slow, percentage

	initialSolution := GenRandomSolution(params.NumDays, blocks.Weights, r)
	solution := SimAnneal(evaluate, findNeighborhood, initialSolution, limitNotImproved, temp, minTemp, alpha, strategy, r)

	if solution.Score.Penalty != 0 {
		t.Error("Expected solution penalty to be 0, got ", solution.Score.Penalty)
	}
}

func TestSimAnneal(t *testing.T) {
	params := paramSets[0]
	SimAnnealTest(t, params)
}

func TestSimAnneal2(t *testing.T) {
	params := paramSets[1]
	SimAnnealTest(t, params)
}

func TestSimAnneal3(t *testing.T) {
	params := paramSets[2]
	SimAnnealTest(t, params)
}

func TestSimAnneal4(t *testing.T) {
	params := paramSets[3]
	SimAnnealTest(t, params)
}
