package sortingoutpost

import (
	"strconv"
	"strings"
)

func uniqueDigitProducts(a []int) int {
	var prod int
	s := 0
	prods := make(map[int]bool)

	for _, n := range a {
		prod = 1

		for _, digit := range strings.Split(strconv.Itoa(n), "") {
			d, _ := strconv.Atoi(digit)
			prod *= d
		}

		if _, ok := prods[prod]; !ok {
			prods[prod] = true
			s++
		}
	}

	return s
}
