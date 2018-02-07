package sortingoutpost

import "sort"

type stringSlice []string

func (s stringSlice) Len() int           { return len(s) }
func (s stringSlice) Less(i, j int) bool { return len(s[i]) < len(s[j]) }
func (s stringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func sortByLength(inputArray []string) []string {
	sort.Sort(stringSlice(inputArray))

	return inputArray
}
