package codefightsgo

import (
	"sort"
)

func SortByHeight(a []int) []int {
	var noTrees []int
	sortedWithTrees := make([]int, 0, len(a))

	// Remove trees
	for _, val := range a {
		if val != -1 {
			noTrees = append(noTrees, val)
		}
	}

	// Sort people
	sort.Ints(noTrees)

	// Put the slice back together again
	for i, noTreesI := 0, 0; i < len(a); i++ {
		if a[i] == -1 {
			sortedWithTrees = append(sortedWithTrees, -1)
		} else {
			sortedWithTrees = append(sortedWithTrees, noTrees[noTreesI])
			noTreesI++
		}
	}

	return sortedWithTrees
}
