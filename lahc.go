package main

import "github.com/jinzhu/copier"

// LAHC implements Late Acceptance Hill Climbing metaheuristic
func LAHC(evaluate Evaluate, findNeighborhood FindNeighborhood, initialSolution []int, maxIter, limitNotImproved, listSize int) Solution {
	count := 0
	bestCost := evaluate(initialSolution)
	solution := Solution{
		Value: initialSolution,
		Score: bestCost,
	}
	bestSolution := solution

	lahcList := make([]Solution, listSize)
	for i := 0; i < listSize; i++ {
		copier.Copy(&lahcList[i], &solution)
	}

	notImprovedCounter := 0

	for count <= maxIter {
		v := count % listSize
		neighbors := findNeighborhood(solution.Value)
		nextSolution := FindBestSolution(neighbors, evaluate)
		if nextSolution.Score.Total < lahcList[v].Score.Total ||
			nextSolution.Score.Total < solution.Score.Total ||
			nextSolution.Score.Penalty < lahcList[v].Score.Penalty ||
			nextSolution.Score.Penalty < solution.Score.Penalty {
			solution = nextSolution
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

		lahcList[v] = solution
		count++
	}
	return bestSolution
}
