package regularhell

import "regexp"

func eyeRhyme(pairOfLines string) bool {
	re := regexp.MustCompile("(.{3})\\\t.*(.{3})$")
	match := re.FindStringSubmatch(pairOfLines)
	return match[1] == match[2]
}
