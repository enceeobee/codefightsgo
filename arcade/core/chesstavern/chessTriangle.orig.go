package chesstavern

// import (
// 	"math"
// )

// type piece struct {
// 	// name string
// 	row int
// 	col int
// }

// // Ok, fuck it, we're brute forcing it. God fucking damn this shitty exercise.
// func chessTriangle(rows, cols int) int {
// 	var totalTriangles int

// 	if rows < 3 && cols < 3 {
// 		return totalTriangles
// 	}

// 	// So maybe we put each piece into every square in the board and see if we can
// 	// build a triangle from there.

// 	/*
// 		Like, put the knight anywhere (like, on every square)
// 		Then put bish/rook on any available knight-move
// 		Then see if we can put the other piece in the second piece's attack zone (without breaking the triangle)
// 	*/

// 	// Move each piece right until the end of the board, then down

// 	return totalTriangles
// }

// func (p *piece) setPiece(row, col int) {
// 	p.row = row
// 	p.col = col
// }

// func (p *piece) moveUp(totalRows, totalCols int) bool {
// 	if p.row > 0 {
// 		p.row--
// 		return true
// 	}
// 	return false
// }

// func (p *piece) moveRight(totalRows, totalCols int) bool {
// 	if p.col < totalRows-1 {
// 		p.col++
// 		return true
// 	}
// 	return false
// }

// func (p *piece) moveDown(totalRows, totalCols int) bool {
// 	if p.row < totalRows-1 {
// 		p.row--
// 		return true
// 	}
// 	return false
// }

// func (p *piece) moveLeft(totalRows, totalCols int) bool {
// 	if p.col > 0 {
// 		p.col--
// 		return true
// 	}
// 	return false
// }

// func isTriangle(bishop, knight, rook piece, totalRows, totalCol int) bool {
// 	// var isBishAlive, isKnightAlive, isRookAlive bool
// 	// var bishIsGood, knightIsGood, rookIsGood bool
// 	isBishAlive, isKnightAlive, isRookAlive := true, true, true

// 	// Bishop takes exactly one piece
// 	// Bishop can take a piece if: diff in rows == diff in cols
// 	var colDiff, rowDiff int
// 	colDiff = bishop.col - knight.col
// 	rowDiff = bishop.row - knight.row
// 	bishopTakesKnight := math.Abs(float64(colDiff)) == math.Abs(float64(rowDiff))

// 	colDiff = bishop.col - rook.col
// 	rowDiff = bishop.row - rook.row
// 	bishopTakesRook := math.Abs(float64(colDiff)) == math.Abs(float64(rowDiff))

// 	if bishopTakesKnight && bishopTakesRook {
// 		return false
// 	} else if bishopTakesKnight && !bishopTakesRook {
// 		isKnightAlive = false
// 	} else if !bishopTakesKnight && bishopTakesRook {
// 		isRookAlive = false
// 	}

// 	// Rook takes another
// 	// Rook can take a piece if: diff in rows = 0, or dif in cols = 0
// 	var rookTakesBishop, rookTakesKnight bool
// 	colDiff = rook.col - bishop.col
// 	rowDiff = rook.row - bishop.row
// 	rookTakesBishop = colDiff == 0 || rowDiff == 0

// 	if isKnightAlive {
// 		colDiff = rook.col - knight.col
// 		rowDiff = rook.row - knight.row
// 		rookTakesKnight = colDiff == 0 || rowDiff == 0
// 	}

// 	if rookTakesBishop && rookTakesKnight {
// 		return false
// 	} else if rookTakesBishop && !rookTakesKnight {
// 		isBishAlive = false
// 	} else if !rookTakesBishop && rookTakesKnight {
// 		isKnightAlive = false
// 	}

// 	// Knight takes the other
// 	var alivePiece piece
// 	if isBishAlive {
// 		alivePiece = bishop
// 	} else if isRookAlive {
// 		alivePiece = rook
// 	} else {
// 		// Means bishop took rook, and rook took bishop (uhoh)
// 		return false
// 	}
// 	// Up 1, right 2
// 	if knight.row == alivePiece.row-1 && knight.col == alivePiece.col+2 {
// 		// Hmm, is this right?
// 		return true
// 	}
// 	// Up 2, right 1
// 	if knight.row == alivePiece.row-2 && knight.col == alivePiece.col+1 {
// 		return true
// 	}
// 	// Up 1, left 2
// 	if knight.row == alivePiece.row-1 && knight.col == alivePiece.col-2 {
// 		return true
// 	}
// 	// Up 2, left 1
// 	if knight.row == alivePiece.row-2 && knight.col == alivePiece.col-1 {
// 		return true
// 	}
// 	// Down 1, right 2
// 	if knight.row == alivePiece.row+1 && knight.col == alivePiece.col+2 {
// 		return true
// 	}
// 	// Down 2, right 1
// 	if knight.row == alivePiece.row+2 && knight.col == alivePiece.col+1 {
// 		return true
// 	}
// 	// Down 1, left 2
// 	if knight.row == alivePiece.row+1 && knight.col == alivePiece.col+2 {
// 		return true
// 	}
// 	// Down 2, left 1
// 	if knight.row == alivePiece.row+2 && knight.col == alivePiece.col+1 {
// 		return true
// 	}

// 	// return bishIsGood && knightIsGood && rookIsGood
// 	return false
// }

// /*
// 	Try this approach:
// 	1. Ok, first, determine how many 2x3s can fit in our board.
// 		* in a 3x3, we can fit 4 2x3s. n = 4
// 	3. Multiply that number by 8 (the number of tris for a 2x3)
// 		* 4 * 8 = 32
// 	4. Get ratio of total squares
// 		* 2x3 = 6; 3x3 = 9
// 		* 6/9 = 2/3
// 	5. Divide 32 by ratio
// 		32 / (2/3) = 48

// 	This should work.

// 	... but it doesn't
// */

// // func chessTriangle(n int, m int) int {
// // 	var numHorizontal2x3s, numVertical2x3s int

// // 	if n >= 2 && m >= 3 {
// // 		perRow := m - 2
// // 		validRows := int(math.Max(float64(n-1), 1))
// // 		numHorizontal2x3s = perRow * validRows
// // 	}
// // 	if n >= 3 && m >= 2 {
// // 		perCol := n - 2
// // 		validCols := int(math.Max(float64(m-1), 1))
// // 		numVertical2x3s = perCol * validCols
// // 	}

// // 	num2x3s := numHorizontal2x3s + numVertical2x3s

// // 	fmt.Println("numHorizontal2x3s", numHorizontal2x3s)
// // 	fmt.Println("numVertical2x3s", numVertical2x3s)
// // 	fmt.Println("num2x3s", num2x3s)

// // 	if num2x3s == 0 {
// // 		return 0
// // 	}

// // 	area := n * m
// // 	boardAreaRatio := 6 / float64(area)
// // 	trianglesFrom2x3s := float64(num2x3s) * 8 // 568 for 7x8

// // 	// so the ratio would need to be 0.298319327731092

// // 	// How much of the board is a 2x3?
// // 	// for a 2x3 it's 100% or 1.0
// // 	// for a 3x3 it's 2/3ds or 66.6% or 0.666
// // 	// for a 7x8 it's 6/56 or 3/28 or 10.7% or x0.107142857142857

// // 	return int(trianglesFrom2x3s / boardAreaRatio)
// // }

// /*
// 	7x8
// 	ans = trisFrom23s / boardAreaRatio
// 	1904 = trisFrom23s / 0.10714285714285714
// 	1904 * 0.10714285714285714 = trisFrom23s
// 	204 = trisFrom23s
// 	204 / 8 = 25.5
// */
