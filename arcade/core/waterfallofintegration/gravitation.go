package waterfallofintegration

import "math"

func gravitation(rows []string) []int {
	secToCols := make(map[int][]int)
	fastest := float64(len(rows))

	for c := 0; c < len(rows[0]); c++ {
		hasBricks := false
		sec := 0

		for r := 0; r < len(rows); r++ {
			if rows[r][c] == '#' {
				hasBricks = true
			} else {
				if hasBricks {
					sec++
				}
			}
		}

		if _, ok := secToCols[sec]; !ok {
			secToCols[sec] = []int{}
		}
		secToCols[sec] = append(secToCols[sec], c)
		fastest = math.Min(fastest, float64(sec))
	}

	return secToCols[int(fastest)]
}
