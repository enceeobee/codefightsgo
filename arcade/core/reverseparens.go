package codefightsgo

import (
	"strings"
)

func ReverseParentheses(s string) string {
	if !strings.Contains(s, "(") {
		return s
	}
	openParenIndex := strings.LastIndex(s, "(") + 1
	closeParenIndex := strings.Index(s[openParenIndex:], ")") + openParenIndex
	subString := s[openParenIndex:closeParenIndex]

	runes := []rune(subString)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return ReverseParentheses(strings.Replace(s, "("+subString+")", string(runes), 1))
}
