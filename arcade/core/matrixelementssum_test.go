package codefightsgo

import (
	"testing"
)

type testpair struct {
	in   [][]int
	want int
}

func TestMatrixElementSum(t *testing.T) {
	cases := []testpair{
		{
			[][]int{
				[]int{0, 1, 1, 2},
				[]int{0, 5, 0, 0},
				[]int{2, 0, 3, 3},
			},
			9,
		},
		{
			[][]int{
				[]int{1, 1, 1, 0},
				[]int{0, 5, 0, 1},
				[]int{2, 1, 3, 10},
			},
			9,
		},
		{
			[][]int{
				[]int{1, 1, 1},
				[]int{2, 2, 2},
				[]int{3, 3, 3},
			},
			18,
		},
		{
			[][]int{
				[]int{0},
			},
			0,
		},
	}

	for _, c := range cases {
		got := MatrixElementsSum(c.in)
		if got != c.want {
			t.Errorf("TestMatrixElementSum(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
