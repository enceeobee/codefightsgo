package springofintegration

func beautifulText(inputString string, l, r int) bool {
	var j int
	isBeaut := false

	for i := l; i <= r; i++ {

		isBeaut = true

		for j = i; j < len(inputString); j += i {

			if string(inputString[j]) != " " {
				isBeaut = false
				break
			}

			// Advance j
			j++

			if isBeaut && len(inputString)-j == i {
				return true
			}
		}
	}

	return false
}
