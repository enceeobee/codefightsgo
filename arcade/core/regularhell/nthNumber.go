package regularhell

import (
	"regexp"
	"strconv"
)

func nthNumber(s string, n int) string {
	re := regexp.MustCompile("(?:(?:\\d[a-zA-Z0-9]*).){" + strconv.Itoa(n-1) + "}\\D*0*(\\d*)")
	return re.FindStringSubmatch(s)[1]
}
