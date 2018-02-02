package codefightsgo

import (
	"reflect"
	"testing"
)

func TestSortByHeight(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
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

	for _, test := range tests {
		actual := SortByHeight(test.in)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("For %d: expected %d; got %d", test.in, test.expected, actual)
		}
	}
}
