package sortingoutpost

import (
	"sort"
	"strconv"
	"strings"
)

type sortNum struct {
	difference int
	index      int
	val        int
}

type sortNums []sortNum

func (sn sortNums) Len() int { return len(sn) }
func (sn sortNums) Less(i, j int) bool {
	return sn[i].difference < sn[j].difference ||
		(sn[i].difference == sn[j].difference && sn[i].index > sn[j].index)
}
func (sn sortNums) Swap(i, j int) { sn[i], sn[j] = sn[j], sn[i] }

func digitDifferenceSort(a []int) []int {
	nums := []sortNum{}

	for i, num := range a {
		s := strings.Split(strconv.Itoa(num), "")
		sort.Strings(s)
		largest, _ := strconv.Atoi(s[len(s)-1])
		smallest, _ := strconv.Atoi(s[0])

		nums = append(nums, sortNum{
			difference: largest - smallest,
			index:      i,
			val:        num,
		})
	}

	sort.Sort(sortNums(nums))

	numsOnly := []int{}

	for _, sn := range nums {
		numsOnly = append(numsOnly, sn.val)
	}

	return numsOnly
}
