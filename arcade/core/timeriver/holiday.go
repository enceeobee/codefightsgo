package timeriver

import (
	"fmt"
	"time"
)

func holiday(x int, weekDay string, month string, yearNumber int) int {
	holidayMonth, _ := time.Parse("January 2006", fmt.Sprintf("%s %d", month, yearNumber))

	for holidayMonth.Weekday().String() != weekDay {
		holidayMonth = holidayMonth.AddDate(0, 0, 1)
	}

	xThDate := holidayMonth.AddDate(0, 0, 7*(x-1))

	if xThDate.Month() == holidayMonth.Month() {
		return xThDate.Day()
	}

	return -1
}
