package timeriver

import (
	"strconv"
)

func validTime(time string) bool {
	hr, err := strconv.Atoi(time[0:2])
	if err != nil {
		return false
	}
	if hr < 0 || hr > 23 {
		return false
	}

	min, err := strconv.Atoi(time[3:])
	if err != nil {
		return false
	}
	if min < 0 || min > 59 {
		return false
	}

	return true
}
