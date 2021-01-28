package main

import "math/rand"

// NextRandom gets next random value within the interval including min and max
func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.Max-ir.Min+1) + ir.Min
}

// TestEq tests slices equality
func TestEq(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// CheckIfSliceInSlices checks if slice "a" is present inside the slice "b"
func CheckIfSliceInSlices(a []int, b [][]int) bool {
	for i := range b {
		if TestEq(a, b[i]) == true {
			return true
		}
	}
	return false
}

// FilterSliceOfSlices filters out slices that are not present in both "a" and "b"
func FilterSliceOfSlices(a [][]int, b [][]int) [][]int {
	var result [][]int

	for i := range a {
		if CheckIfSliceInSlices(a[i], b) == false {
			result = append(result, a[i])
		}
	}

	return result
}

// UniqueSlices gets a unique set of slices
func UniqueSlices(a [][]int) [][]int {
	var result [][]int
	for _, entry := range a {
		if CheckIfSliceInSlices(entry, result) == false {
			result = append(result, entry)
		}
	}

	return result
}
