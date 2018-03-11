package chesstavern

import (
	"fmt"
)

func chessBishopDream(boardSize []int, initPosition []int, initDirection []int, k int) []int {
	dir, pos := make([]int, 2), make([]int, 2)
	copy(dir, initDirection)
	copy(pos, initPosition)

	var xDelay, yDelay int
	var vector string

	seq := [][]int{}
	vectors := make(map[string]bool)

	for i := 0; i < k; i++ {
		xDelay, yDelay = 0, 0
		vector = fmt.Sprintf("%v%v", pos, dir)

		if _, ok := vectors[vector]; !ok {
			vectors[vector] = true

			// This is weird and tricky
			thisPos := make([]int, 2)
			copy(thisPos, pos)

			seq = append(seq, thisPos)
		} else {
			// We've completed a loop, fast forward to end
			return seq[k%len(seq)]
		}

		// Hit top or bottom
		if (pos[0] == 0 && dir[0] < 0) ||
			(pos[0] == boardSize[0]-1 && dir[0] > 0) {
			dir[0] = toggleSign(dir[0])
			yDelay = toggleSign(dir[0])
		}

		// Hit left or right
		if (pos[1] == 0 && dir[1] < 0) ||
			(pos[1] == boardSize[1]-1 && dir[1] > 0) {
			dir[1] = toggleSign(dir[1])
			xDelay = toggleSign(dir[1])
		}

		pos[0] += dir[0] + yDelay
		pos[1] += dir[1] + xDelay
	}

	return pos
}

func toggleSign(n int) int {
	return n - (n * 2)
}
