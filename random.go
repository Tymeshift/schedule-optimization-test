package main

import (
	"math/rand"
)

// Random implements Random Walk
func Random(evaluate Evaluate, initialSolution []int, params ProblemParams, blocks BlocksData, maxIter, limitNotImproved int, r *rand.Rand) Solution {
	count := 0
	bestCost := evaluate(initialSolution)
	solution := Solution{
		Value: initialSolution,
		Score: bestCost,
	}
	bestSolution := solution

	notImprovedCounter := 0

	for count <= maxIter {

		nextSolutionValue := GenRandomSolution(params.NumDays, blocks.Weights, r)
		nextSolutionScore := evaluate(nextSolutionValue)
		solution := Solution{
			Value: nextSolutionValue,
			Score: nextSolutionScore,
		}

		cost := bestSolution.Score.Total - solution.Score.Total
		if cost >= 0 && solution.Score.Penalty <= bestSolution.Score.Penalty {
			notImprovedCounter = 0
			bestSolution = solution
		} else {
			notImprovedCounter++
		}

		if bestSolution.Score.Penalty == 0 && notImprovedCounter >= limitNotImproved {
			return bestSolution
		}

		count++
	}
	return bestSolution
}
