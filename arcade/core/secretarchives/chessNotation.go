package secretarchives

import (
	"strconv"
	"strings"
)

/*
	00 01 02      20 10 00
	10 11 12  ->  21 11 01
	20 21 22      22 12 02

	so:
		row 0 becomes col[length - 1] (aka col[7]) BUT UPSIDE DOWN
			00 = 20
			01 = 10
			02 = 00
		row 1 becomes col[len - 2] BUT UPSIDE DOWN
			10 = 21
			11 = 11
			12 = 01
		row 2 becomes col[len - 3] BUT UPSIDE DOWN
			20 = 22
			21 = 12
			21 = 02
*/

func chessNotation(notation string) string {
	var rotatedNotation string
	board := [][]string{}

	for _, row := range strings.Split(notation, "/") {
		thisRow := []string{}
		rowVals := strings.Split(row, "")

		for _, item := range rowVals {
			n, err := strconv.Atoi(item)

			if err != nil {
				thisRow = append(thisRow, item)
			} else {
				for k := 0; k < n; k++ {
					thisRow = append(thisRow, "x")
				}
			}
		}
		board = append(board, thisRow)
	}

	tempBoard := [][]string{}
	tempBoard = append(tempBoard, make([]string, 8), make([]string, 8), make([]string, 8), make([]string, 8), make([]string, 8), make([]string, 8), make([]string, 8), make([]string, 8))

	rowCounter := 7
	for ri := 0; ri < 8; ri++ {
		rowCounter = 7
		for ci := 0; ci < 8; ci++ {
			tempBoard[ri][ci] = board[rowCounter][ri] // first col, bottom to top
			rowCounter--
		}
	}

	emptyCounter := 0
	for i, row := range tempBoard {
		for _, item := range row {
			if item == "x" {
				emptyCounter++
			} else {
				if emptyCounter != 0 {
					rotatedNotation += strconv.Itoa(emptyCounter)
					emptyCounter = 0
				}
				rotatedNotation += item
			}
		}

		if emptyCounter != 0 {
			rotatedNotation += strconv.Itoa(emptyCounter)
			emptyCounter = 0
		}

		if i < 7 {
			rotatedNotation += "/"
		}
	}

	return rotatedNotation
}
