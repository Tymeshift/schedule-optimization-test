package main

// ProblemParams is a struct with all relevant params for the problem
type ProblemParams struct {
	NumDays    int
	MaxWorking int
	MinWorking int
	MaxOff     int
	MinOff     int
	DaysOff    []int
}

var paramSets = []ProblemParams{{
	NumDays:    28,
	MaxWorking: 5,
	MinWorking: 2,
	MaxOff:     3,
	MinOff:     1,
	DaysOff:    []int{5, 6},
}, {
	NumDays:    28,
	MaxWorking: 5,
	MinWorking: 2,
	MaxOff:     3,
	MinOff:     1,
	DaysOff:    []int{2, 5},
}, {
	NumDays:    112,
	MaxWorking: 5,
	MinWorking: 2,
	MaxOff:     3,
	MinOff:     1,
	DaysOff:    []int{5, 6},
}, {
	NumDays:    112,
	MaxWorking: 4,
	MinWorking: 1,
	MaxOff:     4,
	MinOff:     2,
	DaysOff:    []int{1, 4, 5, 6},
}}

func isInSlice(a int, b []int) bool {
	for _, i := range b {
		if a == i {
			return true
		}
	}
	return false
}

func getSearchParams(params ProblemParams) (Evaluate, FindNeighborhood, BlocksData) {
	// template for mandatory pre-defined day-offs
	var workingDays []int
	for i := 0; i < params.NumDays; i++ {
		if i != 0 && isInSlice(i%7, params.DaysOff) {
			workingDays = append(workingDays, 0)
		} else {
			workingDays = append(workingDays, 1)
		}
	}

	blocks := CreateBlocks(params.MinOff, params.MaxOff, params.MinWorking, params.MaxWorking)
	evaluate := EvaluateFactory(params.NumDays, workingDays, blocks.Weights, blocks.Blocks, params.MinOff, params.MaxOff, params.MinWorking, params.MaxWorking)
	findNeighborhood := FindNeighborhoodFactory(len(blocks.Weights))
	return evaluate, findNeighborhood, blocks
}
