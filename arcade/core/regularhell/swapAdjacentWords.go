package regularhell

import "regexp"

func swapAdjacentWords(s string) string {
	return regexp.MustCompile("(?:([a-zA-Z]+) ([a-zA-Z]+))+").ReplaceAllString(s, "$2 $1")
}
