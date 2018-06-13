package regularhell

import (
	"regexp"
)

func isSubsequence(t string, s string) bool {
	pattern := ""
	for _, ch := range s {
		pattern += "(?:.*)[" + string(ch) + "]"
	}
	re := regexp.MustCompile(pattern)
	return re.MatchString(t)
}
