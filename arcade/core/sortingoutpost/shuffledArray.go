package sortingoutpost

import (
	"sort"
)

func shuffledArray(shuffled []int) []int {
	var sum int

	for i := range shuffled {
		sum = 0
		// Get the sum for all but ai ints
		for ai, an := range shuffled {
			if ai == i {
				continue
			}
			sum += an
		}

		// If this sum is one of the array elements, we've found the sum!
		for ji, jn := range shuffled {
			if sum == jn {
				noSum := append(shuffled[:ji], shuffled[ji+1:]...)
				sort.Ints(noSum)
				return noSum
			}
		}
	}

	sort.Ints(shuffled)
	return shuffled
}
