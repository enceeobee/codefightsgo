package timeriver

import (
	"strconv"
	"time"
)

func dayOfWeek(birthdayDate string) int {
	var date time.Time

	y, _ := strconv.Atoi(birthdayDate[6:])
	m, _ := strconv.Atoi(birthdayDate[0:2])
	d, _ := strconv.Atoi(birthdayDate[3:5])

	bday := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	oWd := bday.Weekday()
	i := 1

	for i < 1000 {
		date = time.Date(y+i, time.Month(m), d, 0, 0, 0, 0, time.UTC)
		i++

		if bday.Day() != date.Day() {
			continue
		}

		if date.Weekday() == oWd {
			return i - 1
		}
	}

	return -1
}
