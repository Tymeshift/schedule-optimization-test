package main

import (
	"math/rand"
	"sort"
)

// FindBestSolution finds the best neighbor
func FindBestSolution(neighbors [][]int, evaluate Evaluate) Solution {
	evaluated := make([]Solution, len(neighbors))
	for i := range neighbors {
		evaluated[i] = Solution{
			Value: neighbors[i],
			Score: evaluate(neighbors[i]),
		}
	}

	sort.Slice(evaluated, func(i, j int) bool {
		return evaluated[i].Score.Penalty < evaluated[j].Score.Penalty
	})

	return evaluated[0]
}

// GetTotalWeight calculates total weight of the solution
func GetTotalWeight(weights []int, solution []int) int {
	totalWeight := 0
	for _, s := range solution {
		totalWeight += weights[s]
	}
	return totalWeight
}

// GenRandomSolution generates random solution, to be used as the initial solution for the search
func GenRandomSolution(maxWeight int, weights []int, r *rand.Rand) []int {
	ir := IntRange{Min: 0, Max: len(weights) - 1}
	curWeight := 0
	var blocks []int
	for maxWeight > curWeight {
		randBlock := ir.NextRandom(r)
		curWeight += weights[randBlock]
		if curWeight <= maxWeight {
			blocks = append(blocks, randBlock)
		}
	}

	return blocks
}

// FindNeighborhoodFactory creates a FindNeighborhood function to be used in a search
func FindNeighborhoodFactory(dataLength int) FindNeighborhood {
	return func(solution []int) [][]int {
		var neighbors [][]int
		for i := range solution {
			newSolution := make([]int, len(solution))
			copy(newSolution, solution)
			if newSolution[i] > 0 {
				newSolution[i]--
				neighbors = append(neighbors, newSolution)
			}

			newSolution = make([]int, len(solution))
			copy(newSolution, solution)
			if newSolution[i] < dataLength-1 {
				newSolution[i]++
				neighbors = append(neighbors, newSolution)
			}

			newSolution = make([]int, len(solution))
			newSolution = append(newSolution[:i], newSolution[i+1:]...)
			neighbors = append(neighbors, newSolution)

			for z := 0; z < dataLength; z++ {
				newSolution = make([]int, len(solution))
				copy(newSolution, solution)
				newSolution = append(newSolution[:i], z)
				newSolution = append(newSolution, newSolution[i:]...)
				neighbors = append(neighbors, newSolution)
			}
		}

		return UniqueSlices(neighbors)
	}
}

// BlocksData contains the blocks and associated parameters that are used to constract the solution. Inspired by knapsack problem
type BlocksData struct {
	Weights []int
	Values  []int
	Blocks  [][]int
}

// CreateBlocks generate all possible blocks based on the problem parameters
func CreateBlocks(minOff int, maxOff int, minWorking int, maxWorking int) BlocksData {
	var weights []int
	var values []int
	var blocks [][]int
	for i := minWorking; i <= maxWorking; i++ {
		for z := minOff; z <= maxOff; z++ {
			weights = append(weights, i+z)
			values = append(values, i)
			blocks = append(blocks, []int{i, z})
		}
	}

	return BlocksData{
		Weights: weights,
		Values:  values,
		Blocks:  blocks,
	}
}

// BlocksToSolution converts block indices to the actual values
func BlocksToSolution(blocks [][]int, solution []int) [][]int {
	var days [][]int
	for i, s := range solution {
		days[i] = blocks[s]
	}

	return days
}
