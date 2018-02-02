package codefightsgo

func AllLongestStrings(inputArray []string) []string {
	var lengths map[int][]string
	lengths = make(map[int][]string)
	maxLen := 0

	for _, word := range inputArray {
		if len(word) > maxLen {
			maxLen = len(word)
		}
		lengths[len(word)] = append(lengths[len(word)], word)
	}

	return lengths[maxLen]
}
