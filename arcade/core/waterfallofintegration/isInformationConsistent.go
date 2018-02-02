package waterfallofintegration

func isInformationConsistent(evidences [][]int) bool {
	var hasGuilty, hasInno bool

	for c := 0; c < len(evidences[0]); c++ {
		hasGuilty, hasInno = false, false

		for r := 0; r < len(evidences); r++ {
			if evidences[r][c] == 1 && !hasGuilty {
				hasGuilty = true
			} else if evidences[r][c] == -1 && !hasInno {
				hasInno = true
			}

			if hasGuilty && hasInno {
				return false
			}
		}
	}
	return true
}
