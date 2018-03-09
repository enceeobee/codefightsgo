package chesstavern

import "math"

func bishopAndPawn(bishop string, pawn string) bool {
	colDiff := math.Abs(float64(bishop[0]) - float64(pawn[0]))
	rowDiff := math.Abs(float64(bishop[1]) - float64(pawn[1]))

	return colDiff == rowDiff
}
