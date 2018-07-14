package secretarchives

import (
	"regexp"
)

func htmlTable(table string, row int, column int) string {
	var cols [][]string

	rowsRegex := regexp.MustCompile(`(?U)<tr>(.*)</tr>`)
	colsRegex := regexp.MustCompile(`(?U)<td>(.*)</td>`)
	rows := rowsRegex.FindAllStringSubmatch(table, -1)

	for rowNum := 1; rowNum < len(rows)+1; rowNum++ {
		if rowNum-1 == row {
			cols = colsRegex.FindAllStringSubmatch(rows[rowNum-1][0], -1)

			if column < len(cols) {
				return cols[column][1]
			} else {
				return "No such cell"
			}
		}
	}

	return "No such cell"
}
