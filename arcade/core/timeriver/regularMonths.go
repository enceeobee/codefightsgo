package timeriver

import (
	"time"
)

func regularMonths(curMonth string) string {
	date, _ := time.Parse("01-2006", curMonth)
	nextMonth := date.AddDate(0, 1, 0)

	for i := 0; i < 100; i++ {
		if nextMonth.Day() == 1 && nextMonth.Weekday() == 1 {
			return nextMonth.Format("01-2006")
		}
		nextMonth = nextMonth.AddDate(0, 1, 0)
	}

	return "Did not find a regular month"
}
