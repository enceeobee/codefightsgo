package springofintegration

import (
	"math"
)

// This is ugly, but it's how my brain worked when solving this puzzle

const filler = '$'

func combs(comb1 string, comb2 string) int {
	minlen := float64(len(comb1) + len(comb2))
	combs := [][]byte{
		[]byte{},
		[]byte{},
	}

	// populate combs
	for i := 0; i < (len(comb1)*2)+len(comb2); i++ {
		if i < len(comb1) {
			combs[0] = append(combs[0], comb1[i])
			combs[1] = append(combs[1], byte(filler))
		} else if i-len(comb1) < len(comb2) {
			combs[0] = append(combs[0], byte(filler))
			combs[1] = append(combs[1], comb2[i-len(comb1)])
		} else {
			combs[0] = append(combs[0], byte(filler))
			combs[1] = append(combs[1], byte(filler))
		}
	}

	for combs[0][len(combs[0])-1] == byte(filler) {
		// If no conflict, get len, then recalculate minlen
		if getConflict(combs) == -1 {
			minlen = math.Min(minlen, getLen(combs))
		}

		// Rotate
		newC := append([]byte{byte(filler)}, combs[0]...)
		newC = newC[0 : len(newC)-1]
		copy(combs[0], newC)
	}

	return int(minlen)
}

func getConflict(combs [][]byte) int {
	for i := 0; i < len(combs[0]); i++ {
		if combs[0][i] == byte('*') && combs[1][i] == byte('*') {
			return i
		}
	}

	return -1
}

func getLen(combs [][]byte) float64 {
	end := len(combs[0]) - 1
	start := -1

	for i := 0; i < len(combs[0]); i++ {
		if (combs[0][i] != byte(filler) || combs[1][i] != byte(filler)) && start < 0 {
			start = i
		}
		if combs[0][i] == byte(filler) && combs[1][i] == byte(filler) && start > -1 {
			end = i
			break
		}
	}

	return float64(end - start)
}
