package regularhell

import (
	"regexp"
	"strconv"
)

func nthNumber(s string, n int) string {
	re := regexp.MustCompile("(?:\\d+\\D+){" + strconv.Itoa(n-1) + "}\\D*0*(\\d*)")
	return re.FindStringSubmatch(s)[1]
}
