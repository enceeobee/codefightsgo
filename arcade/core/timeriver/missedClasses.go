package timeriver

import (
	"fmt"
	"time"
)

/*
 Notes:
	 - daysOfTheWeek are 1-based; 1 = Mon, 2 = Tues, ... 7 = Sun
		 - Sunday will differ, but the rest will match

*/
// Mon Jan 2 15:04:05 MST 2006

func missedClasses(year int, daysOfTheWeek []int, holidays []string) int {
	missedClassesCount := 0

	daysOfTheWeekMap := map[int]bool{}
	for _, dayOfTheWeek := range daysOfTheWeek {
		daysOfTheWeekMap[dayOfTheWeek] = true

		if dayOfTheWeek == 7 {
			daysOfTheWeekMap[0] = true
		}
	}

	var parseValue string
	var holiday time.Time

	for i := range holidays {
		parseValue = fmt.Sprintf("%d-%s-%s", year, holidays[i][0:2], holidays[i][3:])
		holiday, _ = time.Parse("2006-01-02", parseValue)

		if int(holiday.Month()) < 6 {
			holiday = holiday.AddDate(1, 0, 0)
		}

		if _, ok := daysOfTheWeekMap[int(holiday.Weekday())]; ok {
			missedClassesCount++
		}
	}

	return missedClassesCount
}
