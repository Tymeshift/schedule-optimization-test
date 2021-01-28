package main

// TabuSearch - performs the search using the tabu search technique
func TabuSearch(evaluate Evaluate, findNeighborhood FindNeighborhood, initialSolution []int, maxIter, limitNotImproved, tabuSize int) Solution {
	count := 0
	bestCost := evaluate(initialSolution)
	solution := Solution{
		Value: initialSolution,
		Score: bestCost,
	}
	bestSolution := solution
	var tabuList [][]int

	notImprovedCounter := 0

	for count <= maxIter {
		// get all of the neighbors
		neighbors := findNeighborhood(solution.Value)
		filteredNeighbors := FilterSliceOfSlices(neighbors, tabuList)

		if len(filteredNeighbors) > 0 {
			solution = FindBestSolution(filteredNeighbors, evaluate)
			// get the cost between the two solutions
			cost := bestSolution.Score.Total - solution.Score.Total
			// if the new solution is better,
			// update the current solution with the new solution
			if cost >= 0 && solution.Score.Penalty <= bestSolution.Score.Penalty {
				notImprovedCounter = -1
				bestSolution = solution
			}
			notImprovedCounter++

			if bestSolution.Score.Penalty == 0 && notImprovedCounter >= limitNotImproved {
				return bestSolution
			}

			tabuVal := make([]int, len(solution.Value))
			copy(tabuVal, solution.Value)
			tabuList = append(tabuList, tabuVal)

			if len(tabuList) > tabuSize {
				tabuList = tabuList[1:]
			}

		}

		count++
	}

	return bestSolution
}
