package sortingoutpost

import (
	"reflect"
	"testing"
)

func TestSortByHeight(t *testing.T) {
	tests := []struct {
		a []int
		x []int
	}{
		{
			[]int{-1, 150, 190, 170, -1, -1, 160, 180},
			[]int{-1, 150, 160, 170, -1, -1, 180, 190},
		},
		{
			[]int{-1, -1, -1, -1, -1},
			[]int{-1, -1, -1, -1, -1},
		},
		{
			[]int{4, 2, 9, 11, 2, 16},
			[]int{2, 2, 4, 9, 11, 16},
		},
	}

	for i, test := range tests {
		actual := sortByHeight(test.a)
		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("Bonked %d - Wanted %v; Got %v", i+1, test.x, actual)
		}
	}
}
