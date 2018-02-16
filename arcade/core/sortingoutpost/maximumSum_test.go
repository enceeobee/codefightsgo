package sortingoutpost

import (
	"testing"
)

func TestMaximumSum(t *testing.T) {
	tests := []struct {
		a []int
		q [][]int
		x int
	}{
		// 1

		// map[0:1 1:3 2:3 3:2 4:1]
		// [2, 9, 7, 4, 4] = 62 - this is the solution
		// 2 + (9*3) + (7*3) + (4*2) + 4 =
		// 2 + 27 + 21 + 8 + 4 = 62
		{
			[]int{9, 7, 2, 4, 4},
			[][]int{
				[]int{1, 3},
				[]int{1, 4},
				[]int{0, 2},
			},
			62,
		},

		// 2
		{
			[]int{2, 1, 2},
			[][]int{
				[]int{0, 1},
			},
			4,
		},

		// 3
		{
			[]int{5, 3, 2},
			[][]int{
				[]int{0, 0},
				[]int{0, 1},
				[]int{1, 2},
				[]int{0, 2},
			},
			28,
		},

		// 4
		{
			[]int{5, 2, 4, 1, 3},
			[][]int{
				[]int{0, 4},
				[]int{1, 2},
				[]int{1, 2},
			},
			33,
		},

		// 5
		{
			[]int{4, 2, 1, 6, 5, 7, 2, 4},
			[][]int{
				[]int{1, 6},
				[]int{2, 4},
				[]int{3, 6},
				[]int{0, 7},
				[]int{3, 6},
				[]int{4, 4},
				[]int{5, 6},
				[]int{5, 6},
				[]int{0, 1},
				[]int{3, 4},
			},
			162,
		},
	}

	var actual int

	for i, test := range tests {
		actual = maximumSum(test.a, test.q)

		if actual != test.x {
			t.Errorf("%d - Expected %d, got %d", i+1, test.x, actual)
		}
	}
}
