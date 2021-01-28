package main

import (
	"math"
	"math/rand"
)

// SimAnneal implements the simmulated annealing metaheuristic
func SimAnneal(evaluate Evaluate, findNeighborhood FindNeighborhood, initialSolution []int, limitNotImproved int, temp, minTemp, alpha float64, strategy string, r *rand.Rand) Solution {
	count := 0
	bestCost := evaluate(initialSolution)
	solution := Solution{
		Value: initialSolution,
		Score: bestCost,
	}
	bestSolution := solution

	notImprovedCounter := 0

	for temp > minTemp {
		neighbors := findNeighborhood(solution.Value)
		if len(neighbors) > 0 {
			ir := IntRange{Min: 0, Max: len(neighbors) - 1}
			randomNeighbor := neighbors[ir.NextRandom(r)]
			neighborScore := evaluate(randomNeighbor)
			nextSolution := Solution{
				Value: randomNeighbor,
				Score: neighborScore,
			}

			if nextSolution.Score.Total < solution.Score.Total || nextSolution.Score.Penalty < solution.Score.Penalty {
				solution = nextSolution
			} else {
				if r.Float64() < AcceptanceProbability(solution.Score.Total, nextSolution.Score.Total, temp) {
					solution = nextSolution
				}
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
		}

		temp = Cooling(temp, strategy, alpha)
		count++
	}

	return bestSolution
}

// Cooling is the function to decrease temperature
func Cooling(temp float64, strategy string, alpha float64) float64 {
	if strategy == "linear" {
		temp -= alpha
	} else if strategy == "slow" {
		temp /= (1 + alpha*temp)
	} else {
		temp *= alpha
	}

	return temp
}

// AcceptanceProbability is the function that will calculate which value that is more likely to be the answer
func AcceptanceProbability(oldCost, newCost int, temp float64) float64 {
	var result float64

	deltaCost := float64(newCost - oldCost)
	result = math.Exp(-1.0 * math.Abs(deltaCost) / temp)

	return result
}
