package codefightsgo

func CommonCharacterCount(s1 string, s2 string) int {
	// var s1Map map[rune]int
	// var s2Map map[rune]int
	s1Map := make(map[rune]int)
	s2Map := make(map[rune]int)
	commonChars := 0

	for _, st1 := range s1 {
		s1Map[st1]++
	}
	for _, st2 := range s2 {
		s2Map[st2]++
	}

	for key, val := range s1Map {
		if s2Map[key] > 0 {
			if s2Map[key] > val {
				commonChars += val
			} else {
				commonChars += s2Map[key]
			}
		}
	}

	return commonChars
}
