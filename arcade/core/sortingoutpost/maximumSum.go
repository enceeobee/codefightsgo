package sortingoutpost

import "sort"

func maximumSum(a []int, q [][]int) int {
	sum := 0
	uses := make([]int, len(a))

	for _, rng := range q {
		for j := rng[0]; j <= rng[1]; j++ {
			uses[j] = uses[j] + 1
		}
	}

	sort.Ints(a)
	sort.Ints(uses)

	for i, n := range a {
		sum += n * uses[i]
	}

	return sum
}
