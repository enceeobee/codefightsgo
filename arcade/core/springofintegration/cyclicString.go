package springofintegration

func cyclicString(s string) int {
	var ssIndex int
	var ss string
	var goodLen bool
	ssLen := 1
	lenS := len(s)

	// Get all possible s substrings
	for ssLen < lenS {
		ssIndex = 0
		for ssIndex < lenS-ssLen {
			goodLen = true
			ss = s[ssIndex : ssIndex+ssLen]
			compareI := 0
			for i := ssLen; i < lenS; i++ {
				if ss[compareI] != s[i] {
					goodLen = false
					break
				}
				if compareI < ssLen-1 {
					compareI++
				}
			}

			if goodLen {
				return ssLen
			}

			ssIndex++
			break
		}
		ssLen++
	}

	return ssLen
}
