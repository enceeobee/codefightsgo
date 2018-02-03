package waterfallofintegration

import (
	"regexp"
	"strings"
)

func correctNonogram(size int, nonogramField [][]string) bool {
	gridStart := (size + 1) / 2

	for c := gridStart; c < len(nonogramField[0]); c++ {
		// Build row and col strings
		row := strings.Join(nonogramField[c][gridStart:], "")
		col := ""
		rowRegex := "^[.]*"
		colRegex := "^[.]*"

		// Row regex
		for _, val := range nonogramField[c][:gridStart] {
			if val != "-" {
				rowRegex += "#{" + val + "}[.]+"
			}
		}

		// Column regex
		for r := 0; r < gridStart; r++ {
			if nonogramField[r][c] != "-" {
				colRegex += "#{" + nonogramField[r][c] + "}[.]+"
			}
		}

		// Replace the final "+" with "*", and add "$"
		rowRegex = rowRegex[:len(rowRegex)-1] + "*$"
		colRegex = colRegex[:len(colRegex)-1] + "*$"

		for r := gridStart; r < len(nonogramField); r++ {
			col += nonogramField[r][c]
		}
		if rowmatch, _ := regexp.MatchString(rowRegex, row); !rowmatch {
			return false
		}

		if colmatch, _ := regexp.MatchString(colRegex, col); !colmatch {
			return false
		}
	}

	return true
}
