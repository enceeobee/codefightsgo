package springofintegration

import (
	"math"
)

/**
create a map of seconds to position
{
	1: {
		4.5: 1
		6: 1
		8: 1
	}
}
*/

func runnersMeetings(startPosition, speed []int) int {
	var numRunners float64
	all := len(startPosition)
	mc := -1.0
	s := 1
	sToPos := make(map[int]map[float64]int)

	// TODO - Set this condition
	for s < 100000 {
		sToPos[s] = make(map[float64]int)
		for i, pos := range startPosition {
			curPos := (float64(s) * (float64(speed[i]) / 60)) + float64(pos)
			_, ok := sToPos[s][curPos]
			if !ok {
				sToPos[s][curPos] = 0
			}
			sToPos[s][curPos]++
		}

		if len(sToPos[s]) < all {
			for pos := range sToPos[s] {
				numRunners = math.Max(numRunners, float64(sToPos[s][pos]))
			}

			// Shortcut return
			if int(numRunners) == all {
				return all
			}

			mc = math.Max(float64(mc), numRunners)
		}

		s++
	}

	return int(mc)
}
