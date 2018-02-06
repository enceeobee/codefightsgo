package sortingoutpost

import (
	"sort"
)

func sortByHeight(a []int) []int {
	notNeg := []int{}
	sorted := []int{}

	for _, i := range a {
		if i != -1 {
			notNeg = append(notNeg, i)
		}
	}

	sort.Ints(notNeg)

	i := 0
	for _, val := range a {
		if val != -1 {
			sorted = append(sorted, notNeg[i])
			i++
		} else {
			sorted = append(sorted, -1)
		}
	}

	return sorted
}
