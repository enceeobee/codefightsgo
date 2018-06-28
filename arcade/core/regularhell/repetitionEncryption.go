package regularhell

import (
	"regexp"
	"strings"
)

func repetitionEncryption(letter string) int {
	pattern := regexp.MustCompile(`[^a-zA-Z]+`)
	words := pattern.Split(letter, -1)
	res := 0
	for i := range words {
		if i < len(words)-1 && strings.ToLower(words[i]) == strings.ToLower(words[i+1]) {
			res++
		}
	}
	return res
}
