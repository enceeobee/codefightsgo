package chesstavern

import (
	"strconv"
	"strings"
)

func whoseTurn(p string) bool {
	// Looks like we can check the color
	// Num light - whose turn
	/*
		0 - white
		1 - black
		2 - white
		3 - black
		4 - white
	*/
	numWhite := 0

	for _, cell := range strings.Split(p, ";") {
		if isWhite(cell) {
			numWhite++
		}
	}

	return numWhite%2 == 0
}

func isWhite(cell string) bool {
	// Odd rows: b,d,f,h
	row, _ := strconv.Atoi(string(cell[1]))

	if row%2 == 0 {
		if cell[0] == 'a' ||
			cell[0] == 'c' ||
			cell[0] == 'e' ||
			cell[0] == 'g' {
			return true
		}
	} else {
		if cell[0] == 'b' ||
			cell[0] == 'd' ||
			cell[0] == 'f' ||
			cell[0] == 'h' {
			return true
		}
	}

	return false
}
