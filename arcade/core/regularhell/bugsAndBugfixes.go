package regularhell

import (
	"regexp"
	"strconv"
)

/*
	`formula` structure:
		0: full match
		1: number of rolls
		2: dieType, aka number of faces on the die
		3: +|- sign and number
*/
func bugsAndBugfixes(rules string) int {
	pattern := regexp.MustCompile(`(\d*)d(\d+)([+|-]\d+)?`)
	formulas := pattern.FindAllStringSubmatch(rules, -1)

	res := 0
	for _, formula := range formulas {
		rolls := 0
		if formula[1] != "" {
			rolls, _ = strconv.Atoi(formula[1])
		} else {
			rolls = 1
		}
		dieType, _ := strconv.Atoi(formula[2])
		formulaMax := rolls * dieType

		if formula[3] != "" {
			b, _ := strconv.Atoi(formula[3][1:len(formula[3])])
			if formula[3][0] == '-' {
				formulaMax = formulaMax - b
			} else {
				formulaMax = formulaMax + b
			}
		}
		res += formulaMax
	}
	return res
}
