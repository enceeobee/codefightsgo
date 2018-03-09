package chesstavern

import (
	"math"
	"sort"
)

func bishopDiagonal(bishop1 string, bishop2 string) []string {
	colDiff := math.Abs(float64(bishop1[0]) - float64(bishop2[0]))
	rowDiff := math.Abs(float64(bishop1[1]) - float64(bishop2[1]))

	if colDiff == rowDiff {
		moveB1Up := bishop1[1] > bishop2[1]
		moveB1Left := bishop1[0] < bishop2[0]

		if moveB1Up && moveB1Left {
			bishop1 = moveUpAndLeft(bishop1)
			bishop2 = moveDownAndRight(bishop2)
		} else if !moveB1Up && moveB1Left {
			bishop1 = moveDownAndLeft(bishop1)
			bishop2 = moveUpAndRight(bishop2)
		} else if moveB1Up && !moveB1Left {
			bishop1 = moveUpAndRight(bishop1)
			bishop2 = moveDownAndLeft(bishop2)
		} else if !moveB1Up && !moveB1Left {
			bishop1 = moveDownAndRight(bishop1)
			bishop2 = moveUpAndLeft(bishop2)
		}
	}

	bs := []string{bishop1, bishop2}
	sort.Strings(bs)
	return bs
}

func moveUpAndLeft(bish string) string {
	for bish[0] > 'a' && bish[1] < '8' {
		col := bish[0] - 1
		row := bish[1] + 1
		bish = string([]byte{col, row})
	}

	return bish
}

func moveDownAndLeft(bish string) string {
	for bish[0] > 'a' && bish[1] > '1' {
		col := bish[0] - 1
		row := bish[1] - 1
		bish = string([]byte{col, row})
	}

	return bish
}

func moveUpAndRight(bish string) string {
	for bish[0] < 'h' && bish[1] < '8' {
		col := bish[0] + 1
		row := bish[1] + 1
		bish = string([]byte{col, row})
	}

	return bish
}

func moveDownAndRight(bish string) string {
	for bish[0] < 'h' && bish[1] > '1' {
		col := bish[0] + 1
		row := bish[1] - 1
		bish = string([]byte{col, row})
	}

	return bish
}
