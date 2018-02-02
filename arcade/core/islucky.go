package codefightsgo

import (
	"strconv"
)

func IsLucky(n int) bool {
	i, firstHalfSum, lastHalfSum, numLen := 0, 0, 0, len(strconv.Itoa(n))
	nLocal := n

	for i < numLen {
		if i < numLen/2 {
			firstHalfSum += nLocal % 10
		} else {
			lastHalfSum += nLocal % 10
		}
		nLocal /= 10
		i++
	}

	return firstHalfSum == lastHalfSum
}

/* func IsLucky(n int) bool {
	firstHalfSum, lastHalfSum, nStr := 0, 0, strconv.Itoa(n)

	fmt.Println("nStr", nStr, "nStr[0]", nStr[0])

	for i := 0; i < len(nStr)/2; i++ {

		fmt.Println("nStr[i]", nStr[i])

		firstHalfSum += int(nStr[i])
		lastHalfSum += int(nStr[len(nStr)-i])
	}

	return firstHalfSum == lastHalfSum
} */
