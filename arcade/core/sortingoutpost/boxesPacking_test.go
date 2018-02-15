package sortingoutpost

import "testing"

func TestBoxesPacking(t *testing.T) {
	tests := []struct {
		l []int
		w []int
		h []int
		x bool
	}{
		// 1
		{
			[]int{1, 3, 2},
			[]int{1, 3, 2},
			[]int{1, 3, 2},
			true,
		},

		// 2
		{
			[]int{1, 1},
			[]int{1, 1},
			[]int{1, 1},
			false,
		},

		// 3
		{
			[]int{3, 1, 2},
			[]int{3, 1, 2},
			[]int{3, 2, 1},
			false,
		},

		// 4
		{
			[]int{2},
			[]int{3},
			[]int{4},
			true,
		},

		// 5
		{
			[]int{5, 7, 4, 1, 2},
			[]int{4, 10, 3, 1, 4},
			[]int{6, 5, 5, 1, 2},
			true,
		},

		// 6
		{
			[]int{6, 4},
			[]int{5, 3},
			[]int{4, 5},
			true,
		},

		// 7
		{
			[]int{6, 3},
			[]int{5, 4},
			[]int{4, 5},
			true,
		},

		// 8
		{
			[]int{6, 3},
			[]int{5, 5},
			[]int{4, 4},
			true,
		},

		// 9
		{
			[]int{883, 807},
			[]int{772, 887},
			[]int{950, 957},
			true,
		},

		// 10
		{
			[]int{6, 5},
			[]int{5, 3},
			[]int{4, 4},
			true,
		},

		// 11
		{
			[]int{4, 10, 3, 1, 4},
			[]int{5, 7, 4, 1, 2},
			[]int{6, 5, 5, 1, 2},
			true,
		},

		// 12
		{
			[]int{10, 8, 6, 4, 1},
			[]int{7, 7, 6, 3, 2},
			[]int{9, 6, 3, 2, 1},
			true,
		},

		// 13
		{
			[]int{9980, 9984, 9981},
			[]int{9980, 9984, 9983},
			[]int{9981, 9984, 9982},
			true,
		},
	}

	for i, test := range tests {
		actual := boxesPacking(test.l, test.w, test.h)

		if actual != test.x {
			t.Errorf("%d - Wanted %t; got %t", i+1, test.x, actual)
		}
	}
}
