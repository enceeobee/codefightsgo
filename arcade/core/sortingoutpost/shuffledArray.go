package sortingoutpost

import (
	"sort"
)

func shuffledArray(shuffled []int) []int {
	var sum int
	sumIndex := 0

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
				sumIndex = ji

				// Remove the sum, then sort
				noSum := []int{}
				noSum = append(noSum, shuffled[:sumIndex]...)
				noSum = append(noSum, shuffled[sumIndex+1:]...)

				sort.Ints(noSum)
				return noSum
			}
		}
	}

	return shuffled
}
