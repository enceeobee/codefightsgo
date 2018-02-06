package sortingoutpost

import (
	"reflect"
	"testing"
)

func TestShuffledArray(t *testing.T) {
	tests := []struct {
		shuffled []int
		x        []int
	}{
		// 1
		{
			[]int{1, 12, 3, 6, 2},
			[]int{1, 2, 3, 6},
		},
		// 2
		{
			[]int{1, -3, -5, 7, 2},
			[]int{-5, -3, 2, 7},
		},
	}

	for i, test := range tests {
		actual := shuffledArray(test.shuffled)

		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("DOH in test %d; wanted %v; got %v", i+1, test.x, actual)
		}
	}
}
