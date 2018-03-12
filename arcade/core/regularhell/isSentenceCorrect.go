package regularhell

import "regexp"

func isSentenceCorrect(sentence string) bool {
	re := regexp.MustCompile(`^[A-Z]([^\.\!\?]*)[\.\!\?]$`)
	return re.MatchString(sentence)
}
