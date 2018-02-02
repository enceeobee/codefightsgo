package waterfallofintegration

import "testing"

// [][]int -> bool
func TestIsInformationConsistent(t *testing.T) {
	tests := []struct {
		e [][]int
		x bool
	}{
		{
			[][]int{
				[]int{0, 1, 0, 1},
				[]int{-1, 1, 0, 0},
				[]int{-1, 0, 0, 1},
			},
			true,
		},
		// 2
		{
			[][]int{
				[]int{1, 0},
				[]int{-1, 0},
				[]int{1, 1},
				[]int{1, 1},
			},
			false,
		},
		// 3
		{
			[][]int{
				[]int{1, -1, 0, 1},
				[]int{1, -1, 0, -1},
			},
			false,
		},
		// 4
		{
			[][]int{
				[]int{0, 0, -1},
				[]int{-1, 1, -1},
				[]int{-1, 1, 0},
				[]int{0, 1, 0},
			},
			true,
		},
	}

	for i, test := range tests {
		actual := isInformationConsistent(test.e)
		if actual != test.x {
			t.Errorf("%d - Wanted %t; Got %t", i+1, test.x, actual)
		}
	}
}
