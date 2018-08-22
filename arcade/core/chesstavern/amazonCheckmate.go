package chesstavern

import (
	"fmt"
	"math"
	"strconv"
)

type acPiece struct {
	row int
	col int
}

func amazonCheckmate(king, amazon string) []int {
	var checkmates, checks, stalemates, validMoves int
	captures := [][]bool{}
	amazonPiece := acPiece{
		col: getCol(amazon),
		row: getRow(amazon),
	}
	kingPiece := acPiece{
		col: getCol(king),
		row: getRow(king),
	}
	var blackKingPiece acPiece

	for row := 0; row < 8; row++ {
		captures = append(captures, []bool{})
		for col := 0; col < 8; col++ {
			blackKingPiece = acPiece{row: row, col: col}
			captures[row] = append(captures[row], isCellCaptured(amazonPiece, kingPiece, blackKingPiece))
		}
	}

	// This is only for debugging
	printCaptures(captures)

	for row, vals := range captures {
		for col := range vals {
			if (amazonPiece.row == row && amazonPiece.col == col) || (kingPiece.row == row && kingPiece.col == col) {
				fmt.Printf("Either amazon or king are on [%d,%d]\n", row, col)
				continue
			}

			blackKingPiece = acPiece{row: row, col: col}
			if canKingCapture(kingPiece, blackKingPiece) {
				continue
			}
			// black's king is under amazon's attack and it cannot make a valid move
			if canAmazonCapture(amazonPiece, kingPiece, blackKingPiece) && !hasValidKingMove(row, col, captures) {
				checkmates++
			}
			// black's king is under amazon's attack but it can reach a safe square in one move
			if canAmazonCapture(amazonPiece, kingPiece, blackKingPiece) && hasValidKingMove(row, col, captures) {
				checks++
			}
			// black's king is on a safe square but it cannot make a valid move
			if !canAmazonCapture(amazonPiece, kingPiece, blackKingPiece) && !hasValidKingMove(row, col, captures) {
				stalemates++
			}
			// black's king is on a safe square and it can make a valid move
			if !canAmazonCapture(amazonPiece, kingPiece, blackKingPiece) && hasValidKingMove(row, col, captures) {
				validMoves++
			}
		}
	}

	return []int{checkmates, checks, stalemates, validMoves}
}

func isCellCaptured(amazon, king, blackKing acPiece) bool {
	// Amazon's cell is safe, unless it's adjacent to white king
	if blackKing.row == amazon.row && blackKing.col == amazon.col {
		return canKingCapture(king, amazon)
	}

	return canAmazonCapture(amazon, king, blackKing) || canKingCapture(king, blackKing)
}

func canAmazonCapture(amazon, king, blackKing acPiece) bool {
	amazToBlkKingRowDiff := math.Abs(float64(amazon.row - blackKing.row))
	amazToBlkKingColDiff := math.Abs(float64(amazon.col - blackKing.col))

	if canCaptureKnightStyle(amazon, blackKing) {
		return true
	}

	if amazon.row == blackKing.row {
		return canCaptureSameRow(amazon, king, blackKing)
	}

	if amazon.col == blackKing.col {
		return canCaptureSameCol(amazon, king, blackKing)
	}

	if amazToBlkKingRowDiff == amazToBlkKingColDiff {
		return canCaptureDiagonal(amazon, king, blackKing)
	}

	return false
	// return canCaptureKnightStyle(amazon, blackKing) ||
	// 	canCaptureSameRow(amazon, king, blackKing) ||
	// 	canCaptureSameCol(amazon, king, blackKing) ||
	// 	canCaptureDiagonal(amazon, king, blackKing)
}

func canKingCapture(king, otherPiece acPiece) bool {
	rowDiff := math.Abs(float64(king.row - otherPiece.row))
	colDiff := math.Abs(float64(king.col - otherPiece.col))

	return rowDiff < 2 && colDiff < 2
}

func canCaptureKnightStyle(amazon, blackKing acPiece) bool {
	var isCaptured bool

	isCaptured = (blackKing.row == amazon.row-2 && blackKing.col == amazon.col+1) ||
		(blackKing.row == amazon.row-2 && blackKing.col == amazon.col-1) ||
		(blackKing.row == amazon.row-1 && blackKing.col == amazon.col+2) ||
		(blackKing.row == amazon.row-1 && blackKing.col == amazon.col-2) ||
		(blackKing.row == amazon.row+2 && blackKing.col == amazon.col+1) ||
		(blackKing.row == amazon.row+2 && blackKing.col == amazon.col-1) ||
		(blackKing.row == amazon.row+1 && blackKing.col == amazon.col+2) ||
		(blackKing.row == amazon.row+1 && blackKing.col == amazon.col-2)

	return isCaptured
}

func canCaptureSameRow(amazon, king, blackKing acPiece) bool {
	if king.row != amazon.row {
		return true
	}

	var canCaptureRight, canCaptureLeft bool

	// Search right
	for col := amazon.col; col < 8; col++ {
		if col == king.col {
			canCaptureRight = false
			break
		}
		if col == blackKing.col {
			canCaptureRight = true
			break
		}
	}

	// Search left
	for col := amazon.col; col >= 0; col-- {
		if col == king.col {
			canCaptureLeft = false
			break
		}
		if col == blackKing.col {
			canCaptureLeft = true
			break
		}
	}

	return canCaptureRight || canCaptureLeft
}

func canCaptureSameCol(amazon, king, blackKing acPiece) bool {
	if king.col != amazon.col {
		return true
	}

	var canCaptureUp, canCaptureDown bool

	// Search up
	for row := amazon.row; row >= 0; row-- {
		if row == king.row {
			canCaptureUp = false
			break
		}
		if row == blackKing.row {
			canCaptureUp = true
			break
		}
	}

	// Search down
	for row := amazon.row; row < 8; row++ {
		if row == king.row {
			canCaptureDown = false
			break
		}
		if row == blackKing.row {
			canCaptureDown = true
			break
		}
	}

	return canCaptureUp || canCaptureDown
}

func canCaptureDiagonal(amazon, king, blackKing acPiece) bool {
	amazToBlkKingRowDiff := math.Abs(float64(amazon.row - blackKing.row))
	amazToWhiteKingRowDiff := math.Abs(float64(amazon.row - king.row))
	amazToWhiteKingColDiff := math.Abs(float64(amazon.col - king.col))

	// Capture if white king is not diagonal to amazon
	if amazToWhiteKingRowDiff != amazToWhiteKingColDiff {
		return true
	}
	// Capture if black king is closer than white king, regardless of direction
	if amazToBlkKingRowDiff <= amazToWhiteKingRowDiff {
		return true
	}

	// White king is closer than black king;
	// determine if both kings are in the same direction
	// If so, return false
	isWhiteKingAbove := king.row < amazon.row
	isWhiteKingRight := king.col > amazon.col
	isWhiteKingBelow := king.row > amazon.row
	isWhiteKingLeft := king.col < amazon.col

	isBlackKingAbove := blackKing.row < amazon.row
	isBlackKingRight := blackKing.col > amazon.col
	isBlackKingBelow := blackKing.row > amazon.row
	isBlackKingLeft := blackKing.col < amazon.col

	if (isWhiteKingAbove && isWhiteKingRight && isBlackKingAbove && isBlackKingRight) ||
		(isWhiteKingBelow && isWhiteKingRight && isBlackKingBelow && isBlackKingRight) ||
		(isWhiteKingBelow && isWhiteKingLeft && isBlackKingBelow && isBlackKingLeft) ||
		(isWhiteKingAbove && isWhiteKingLeft && isBlackKingAbove && isBlackKingLeft) {
		return false
	}

	return true
}

func getCol(pos string) int {
	cols := map[byte]int{
		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
		'h': 7,
	}

	return cols[pos[0]]
}

func getRow(pos string) int {
	row, _ := strconv.Atoi(string(pos[1:]))
	return 8 - row
}

func hasValidKingMove(row, col int, captures [][]bool) bool {
	return (row > 0 && !captures[row-1][col]) ||
		(row > 0 && col < 7 && !captures[row-1][col+1]) ||
		(col < 7 && !captures[row][col+1]) ||
		(row < 7 && col < 7 && !captures[row+1][col+1]) ||
		(row < 7 && !captures[row+1][col]) ||
		(row < 7 && col > 0 && !captures[row+1][col-1]) ||
		(col > 0 && !captures[row][col-1]) ||
		(row > 0 && col > 0 && !captures[row-1][col-1])
}

func printCaptures(captures [][]bool) {
	fmt.Println("captured cells:")

	for _, row := range captures {
		for _, val := range row {
			if val {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println("")
	}
}
