package sortingoutpost

import (
	"testing"
)

func TestRowsRearranging(t *testing.T) {
	var actual bool

	tests := []struct {
		matrix [][]int
		x      bool
	}{
		// 1
		{
			[][]int{
				[]int{2, 7, 1},
				[]int{0, 2, 0},
				[]int{1, 3, 1},
			},
			false,
		},

		// 2
		{
			[][]int{
				[]int{6, 4},
				[]int{2, 2},
				[]int{4, 3},
			},
			true,
		},

		// 3
		{
			[][]int{
				[]int{0, 1},
				[]int{1, 2},
				[]int{2, 3},
				[]int{-1, 4},
			},
			false,
		},

		// 4
		{
			[][]int{
				[]int{2, 2, 2},
				[]int{1, 1, 1},
			},
			true,
		},

		// 5
		{
			[][]int{
				[]int{1, 3, 1},
				[]int{0, 2, 0},
				[]int{1, 7, 2},
			},
			false,
		},

		// 6
		{
			[][]int{
				[]int{3, 34, 2, 5, 2},
				[]int{2, 34, 5, 2, 1},
			},
			false,
		},

		// 7
		{
			[][]int{
				[]int{0},
				[]int{1},
				[]int{2},
				[]int{-1},
			},
			true,
		},
	}

	for i, test := range tests {
		actual = rowsRearranging(test.matrix)

		if actual != test.x {
			t.Errorf("%d - Got %t; Wanted %t", i+1, actual, test.x)
		}
	}
}
