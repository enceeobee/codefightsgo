package sortingoutpost

import (
	"reflect"
	"testing"
)

func TestDigitDifferenceSort(t *testing.T) {
	tests := []struct {
		a []int
		x []int
	}{
		{
			[]int{152, 23, 7, 887, 243},
			[]int{7, 887, 23, 243, 152},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{13098, 1308, 12398, 52433, 213, 424, 213, 243, 12213, 54234, 99487, 81892, 11111, 97897},
			[]int{11111, 97897, 12213, 243, 213, 424, 213, 54234, 52433, 99487, 81892, 12398, 1308, 13098},
		},
	}

	for i, test := range tests {
		actual := digitDifferenceSort(test.a)

		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("BONK %d - Got %v; Wanted %v", i+1, actual, test.x)
		}
	}
}
