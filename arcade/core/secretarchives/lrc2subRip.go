package secretarchives

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func lrc2subRip(lrcLyrics []string, songLength string) []string {
	times := []time.Time{}
	lyrics := []string{}
	rip := []string{}

	var thisTime time.Time
	regex := regexp.MustCompile(`^\[(.+)\]\s?(.*)$`)
	for _, lyric := range lrcLyrics {
		match := regex.FindAllStringSubmatch(lyric, -1)

		timeSegments := strings.Split(match[0][1], ":")
		minuteSegment, _ := strconv.Atoi(timeSegments[0])

		if minuteSegment < 60 {
			thisTime, _ = time.Parse("04:05", match[0][1])
		} else {
			hours := minuteSegment / 60
			minutes := minuteSegment % 60

			thisTime, _ = time.Parse("15:4:5", strconv.Itoa(hours)+":"+strconv.Itoa(minutes)+":"+timeSegments[1])
		}

		times = append(times, thisTime)
		lyrics = append(lyrics, match[0][2])
	}

	lastTime, _ := time.Parse("15:04:05", songLength)
	times = append(times, lastTime)

	var timeLine string
	var nextLyricTime time.Time

	for i, lyric := range lyrics {
		nextLyricTime = times[i+1]

		timeLine = times[i].Format("15:04:05.000") + " --> " + nextLyricTime.Format("15:04:05.000")
		timeLine = strings.Replace(timeLine, ".", ",", 2)

		rip = append(rip, strconv.Itoa(i+1))
		rip = append(rip, timeLine)
		rip = append(rip, lyric)

		if i < len(lyrics)-1 {
			rip = append(rip, "")
		}
	}

	return rip
}
