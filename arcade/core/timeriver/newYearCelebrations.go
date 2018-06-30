package timeriver

import (
	"strconv"
	"time"
)

/*
	Notes:

	Each value in minutes is that number from the initial takeoffTime, not from the last timezone

	Also, `The "00:00" time corresponds to the midnight of 31st of December / 1st of January.`

	---

	We need currentTime, and nextTimezoneTime (when we hit the next timezone)

	{"23:35", []int{60, 90, 140}, 3},

	Before loop:
		curTime = 23:35
		nextTZTime = 23:35 + 60 = 00:35 (celeb++)

	1.
		curTime = nextTZTime - hour = 00:35 - hour = 23:35
		nextTZTime = curTime + (90 - 60) = curTime + 30 = 00:05 (celeb++)
	2.
		curTime = nextTZTime - hour = 00:05 - hour = 23:05
		nextTZTime = curTime + (140 - 90) = curTime + 50 = 23:55
	3.
		curTime = nextTZTime - hour = 23:55 - hour = 22:55
		nextTZTime = N/A
		22:55 < 00:00 (celeb++)

	---
*/

func newYearCelebrations(takeOffTime string, minutes []int) int {
	celebrations := 0
	currentTime, _ := time.Parse("Jan 2 15:04", "Dec 31 "+takeOffTime)
	hourAgo, _ := time.ParseDuration("-1h")

	if len(minutes) == 0 && currentTime.Day() == 31 {
		return 1
	}

	if takeOffTime == "00:00" {
		day, _ := time.ParseDuration("24h")
		currentTime = currentTime.Add(day)
		celebrations++
	}

	minutesUntilNextTimezone, _ := time.ParseDuration(strconv.Itoa(minutes[0]) + "m")
	nextTimeZone := currentTime.Add(minutesUntilNextTimezone)

	if currentTime.Day() == 31 && nextTimeZone.Day() == 1 {
		celebrations++
	}

	for i, minuteAmount := range minutes {
		currentTime = nextTimeZone.Add(hourAgo)

		if i == len(minutes)-1 {
			if currentTime.Day() == 31 || (currentTime.Hour() == 0 && currentTime.Minute() == 0) {
				celebrations++
			}
			continue
		}

		nextDur, _ := time.ParseDuration(strconv.Itoa(minutes[i+1]-minuteAmount) + "m")
		nextTimeZone = currentTime.Add(nextDur)

		if (currentTime.Day() == 31 || (currentTime.Hour() == 0 && currentTime.Minute() == 0)) && nextTimeZone.Day() == 1 {
			celebrations++
		}
	}

	return celebrations
}
