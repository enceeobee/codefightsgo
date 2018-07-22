package chesstavern

import (
	"fmt"
	"math"
)

type piece struct {
	row int
	col int
}

// Ok, fuck it, we're brute forcing it. God fucking damn this shitty exercise.
// So maybe we put each piece into every square in the board and see if we can
// build a triangle from there.

/*
	Like, put the knight anywhere (like, on every square)
	Then put bish/rook on any available knight-move
	Then see if we can put the other piece in the second piece's attack zone (without breaking the triangle)
*/
func chessTriangle(rows, cols int) int {
	var totalTriangles int

	if rows < 3 && cols < 3 {
		return totalTriangles
	}

	calculateTriangles := generateCalculateFunction(rows, cols)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {

			fmt.Println(row, col, calculateTriangles(row, col), "triangles")

			totalTriangles += calculateTriangles(row, col)
		}
	}

	return totalTriangles
}

func generateCalculateFunction(rows, cols int) func(row, col int) int {
	return func(row, col int) int {
		var triangles int
		var bishop, knight, rook piece

		knight = piece{row: row, col: col}

		// Up 2, right 1
		if row-2 >= 0 && col+1 < cols {
			bishop = piece{row: row - 2, col: col + 1}
			rook = piece{row: row - 2, col: col + 1}

			triangles += countRookTris(knight, bishop, rows, cols)
			triangles += countBishopTris(knight, rook, rows, cols)
		}
		// Up 1, right 2
		if row-1 >= 0 && col+2 < cols {
			bishop = piece{row: row - 1, col: col + 2}
			rook = piece{row: row - 1, col: col + 2}

			triangles += countRookTris(knight, bishop, rows, cols)
			triangles += countBishopTris(knight, rook, rows, cols)
		}
		// Up 2, left 1
		if row-2 >= 0 && col-1 >= 0 {
			bishop = piece{row: row - 2, col: col - 1}
			rook = piece{row: row - 2, col: col - 1}

			triangles += countRookTris(knight, bishop, rows, cols)
			triangles += countBishopTris(knight, rook, rows, cols)
		}
		// Up 1, left 2
		if row-1 >= 0 && col-2 >= 0 {
			bishop = piece{row: row - 1, col: col - 2}
			rook = piece{row: row - 1, col: col - 2}

			triangles += countRookTris(knight, bishop, rows, cols)
			triangles += countBishopTris(knight, rook, rows, cols)
		}
		// Down 2, right 1
		if row+2 < rows && col+1 < cols {
			bishop = piece{row: row + 2, col: col + 1}
			rook = piece{row: row + 2, col: col + 1}

			triangles += countRookTris(knight, bishop, rows, cols)
			triangles += countBishopTris(knight, rook, rows, cols)
		}
		// Down 1, right 2
		if row+1 < rows && col+2 < cols {
			bishop = piece{row: row + 1, col: col + 2}
			rook = piece{row: row + 1, col: col + 2}

			triangles += countRookTris(knight, bishop, rows, cols)
			triangles += countBishopTris(knight, rook, rows, cols)
		}
		// Down 2, left 1
		if row+2 < rows && col-1 >= 0 {
			bishop = piece{row: row + 2, col: col - 1}
			rook = piece{row: row + 2, col: col - 1}

			triangles += countRookTris(knight, bishop, rows, cols)
			triangles += countBishopTris(knight, rook, rows, cols)
		}
		// Down 1, left 2
		if row+1 < rows && col-2 >= 0 {
			bishop = piece{row: row + 1, col: col - 2}
			rook = piece{row: row + 1, col: col - 2}

			triangles += countRookTris(knight, bishop, rows, cols)
			triangles += countBishopTris(knight, rook, rows, cols)
		}
		return triangles
	}
}

func countRookTris(knight, bishop piece, totalRows, totalCols int) int {
	var triangles int

	// So far, knight has taken bishop. Now bishop needs to take rook, while rook takes knight
	// We set rook in every available bishop move

	adjustTriangles := func(row, col int, knight piece) int {
		if knight.row == row || knight.col == col {
			return 1
		}
		return 0
	}

	var canMove, canGoUp, canGoDown, canGoLeft, canGoRight bool
	canMove = true
	cellsToMove := 1

	for canMove {
		canGoUp = bishop.row-cellsToMove >= 0
		canGoDown = bishop.row+cellsToMove < totalRows
		canGoRight = bishop.col+cellsToMove < totalCols
		canGoLeft = bishop.col-cellsToMove >= 0

		// Up right
		if canGoUp && canGoRight {
			triangles += adjustTriangles(bishop.row-cellsToMove, bishop.col+cellsToMove, knight)
		}

		// Down right
		if canGoDown && canGoRight {
			triangles += adjustTriangles(bishop.row+cellsToMove, bishop.col+cellsToMove, knight)
		}
		// Down left
		if canGoDown && canGoLeft {
			triangles += adjustTriangles(bishop.row+cellsToMove, bishop.col-cellsToMove, knight)
		}
		// Up left
		if canGoUp && canGoLeft {
			triangles += adjustTriangles(bishop.row-cellsToMove, bishop.col-cellsToMove, knight)
		}

		cellsToMove++
		canMove = (canGoUp && canGoRight) ||
			(canGoDown && canGoRight) ||
			(canGoDown && canGoLeft) ||
			(canGoUp && canGoLeft)
	}

	return triangles
}

func countBishopTris(knight, rook piece, totalRows, totalCols int) int {
	var triangles int

	// So far, knight has taken rook. Now rook needs to take bishop, while bishop takes knight
	// We set bishop in every available rook move

	bishopTakesKnightTriangles := func(row, col int, knight piece) int {
		if math.Abs(float64(knight.row-row)) == math.Abs(float64(knight.col-col)) {
			return 1
		}
		return 0
	}

	var canMove, canGoUp, canGoDown, canGoLeft, canGoRight bool
	canMove = true
	cellsToMove := 1

	for canMove {
		canGoUp = rook.row-cellsToMove >= 0
		canGoDown = rook.row+cellsToMove < totalRows
		canGoRight = rook.col+cellsToMove < totalCols
		canGoLeft = rook.col-cellsToMove >= 0

		// Up - aka bishop is above rook
		if canGoUp {
			triangles += bishopTakesKnightTriangles(rook.row-cellsToMove, rook.col, knight)
		}
		// Right
		if canGoRight {
			triangles += bishopTakesKnightTriangles(rook.row, rook.col+cellsToMove, knight)
		}
		// Down
		if canGoDown {
			triangles += bishopTakesKnightTriangles(rook.row+cellsToMove, rook.col, knight)
		}
		// Left
		if canGoLeft {
			triangles += bishopTakesKnightTriangles(rook.row, rook.col-cellsToMove, knight)
		}

		cellsToMove++
		canMove = canGoUp || canGoDown || canGoLeft || canGoRight
	}

	return triangles
}
