package waterfallofintegration

import "testing"

func TestCorrectNonogram(t *testing.T) {
	tests := []struct {
		size int
		n    [][]string
		x    bool
	}{
		// 1
		{
			5,
			[][]string{
				[]string{"-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "2", "2", "1", "-", "1"},
				[]string{"-", "-", "-", "2", "1", "1", "3", "3"},
				[]string{"-", "3", "1", "#", "#", "#", ".", "#"},
				[]string{"-", "-", "2", "#", "#", ".", ".", "."},
				[]string{"-", "-", "2", ".", ".", ".", "#", "#"},
				[]string{"-", "1", "2", "#", ".", ".", "#", "#"},
				[]string{"-", "-", "5", "#", "#", "#", "#", "#"},
			},
			true,
		},

		// 2
		{
			5,
			[][]string{
				[]string{"-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "1", "-", "-"},
				[]string{"-", "-", "-", "3", "3", "2", "5", "5"},
				[]string{"-", "-", "3", ".", ".", ".", "#", "#"},
				[]string{"-", "2", "2", "#", "#", "#", "#", "#"},
				[]string{"-", "-", "5", "#", "#", "#", "#", "#"},
				[]string{"-", "-", "5", "#", "#", "#", "#", "#"},
				[]string{"-", "-", "2", ".", ".", ".", "#", "#"},
			},
			false,
		},

		// 3
		{
			5,
			[][]string{
				[]string{"-", "-", "-", "-", "-", "1", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "1", "-", "1"},
				[]string{"-", "-", "-", "1", "2", "1", "5", "1"},
				[]string{"-", "-", "3", "#", "#", "#", ".", "."},
				[]string{"-", "2", "1", ".", "#", "#", ".", "#"},
				[]string{"-", "-", "3", ".", ".", "#", "#", "#"},
				[]string{"-", "-", "2", ".", ".", "#", "#", "."},
				[]string{"-", "-", "2", ".", ".", "#", "#", "."},
			},
			false,
		},

		// 4
		{
			5,
			[][]string{
				[]string{"-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "5", "4", "5", "4", "5"},
				[]string{"-", "1", "3", "#", "#", "#", ".", "#"},
				[]string{"-", "-", "5", "#", "#", "#", "#", "#"},
				[]string{"-", "-", "5", "#", "#", "#", "#", "#"},
				[]string{"-", "-", "5", "#", "#", "#", "#", "#"},
				[]string{"-", "3", "1", "#", ".", "#", "#", "#"},
			},
			false,
		},

		// 5
		{
			5,
			[][]string{
				[]string{"-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "-", "2"},
				[]string{"-", "-", "-", "2", "2", "2", "4", "1"},
				[]string{"-", "-", "-", ".", ".", ".", ".", "."},
				[]string{"-", "-", "2", ".", ".", ".", "#", "#"},
				[]string{"-", "-", "3", ".", ".", "#", "#", "#"},
				[]string{"-", "-", "4", "#", "#", "#", "#", "."},
				[]string{"-", "2", "2", ".", "#", "#", "#", "#"},
			},
			false,
		},

		// 6
		{
			5,
			[][]string{
				[]string{"-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "1", "-", "-", "-"},
				[]string{"-", "-", "-", "1", "3", "1", "2", "1"},
				[]string{"-", "-", "1", ".", "#", ".", ".", "."},
				[]string{"-", "-", "-", ".", ".", ".", ".", "."},
				[]string{"-", "-", "1", ".", "#", ".", ".", "."},
				[]string{"-", "1", "2", ".", "#", ".", "#", "#"},
				[]string{"-", "-", "4", "#", "#", "#", "#", "."},
			},
			true,
		},

		// 7
		{
			9,
			[][]string{
				[]string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "-", "4", "3", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "2", "2", "2", "1", "10", "3", "4", "2", "2"},
				[]string{"-", "-", "-", "-", "1", ".", ".", ".", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "-", "5", ".", ".", "#", "#", "#", "#", "#", ".", "."},
				[]string{"-", "-", "-", "-", "7", ".", "#", "#", "#", "#", "#", "#", "#", "."},
				[]string{"-", "-", "-", "-", "9", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
				[]string{"1", "1", "1", "1", "1", "#", ".", "#", ".", "#", ".", "#", ".", "#"},
				[]string{"-", "-", "-", "-", "1", ".", ".", ".", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "-", "1", ".", ".", ".", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "1", "1", ".", ".", "#", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "-", "3", ".", ".", "#", "#", "#", ".", ".", ".", "."},
			},
			false,
		},

		// 8
		{
			9,
			[][]string{
				[]string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "-", "4", "3", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "2", "2", "2", "1", "9", "3", "4", "2", "2"},
				[]string{"-", "-", "-", "-", "1", ".", ".", ".", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "-", "5", ".", ".", "#", "#", "#", "#", "#", ".", "."},
				[]string{"-", "-", "-", "-", "7", ".", "#", "#", "#", "#", "#", "#", "#", "."},
				[]string{"-", "-", "-", "-", "9", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
				[]string{"1", "1", "1", "1", "1", "#", ".", "#", ".", "#", ".", "#", ".", "#"},
				[]string{"-", "-", "-", "-", "1", ".", ".", ".", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "-", "1", ".", ".", ".", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "1", "1", ".", ".", "#", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "-", "3", ".", ".", "#", "#", "#", ".", ".", ".", "."},
			},
			true,
		},

		// 9
		{
			9,
			[][]string{
				[]string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "-", "4", "3", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "2", "2", "2", "1", "9", "3", "4", "2", "2"},
				[]string{"-", "-", "-", "-", "1", ".", ".", ".", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "-", "5", ".", ".", "#", "#", "#", "#", "#", ".", "."},
				[]string{"-", "-", "-", "-", "7", ".", "#", "#", "#", "#", "#", "#", "#", "."},
				[]string{"-", "-", "-", "-", "9", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
				[]string{"1", "1", "1", "1", "1", "#", ".", "#", ".", "#", ".", "#", ".", "#"},
				[]string{"-", "-", "-", "-", "1", ".", ".", ".", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "-", "1", ".", ".", "#", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "1", "1", ".", ".", "#", ".", "#", ".", ".", ".", "."},
				[]string{"-", "-", "-", "-", "3", ".", ".", "#", "#", "#", ".", ".", ".", "."},
			},
			false,
		},

		// 10
		{
			10,
			[][]string{
				[]string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "-", "1", "-", "-", "-", "-", "-", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "1", "2", "1", "-", "-", "1", "1", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "4", "2", "1", "7", "1", "6", "4", "-", "-", "-"},
				[]string{"-", "-", "-", "-", "-", "2", "1", "4", "1", "8", "1", "1", "2", "1", "3"},
				[]string{"-", "-", "-", "3", "3", "#", "#", "#", ".", "#", "#", "#", ".", ".", "."},
				[]string{"-", "-", "-", "-", "1", ".", ".", ".", "#", ".", ".", ".", ".", ".", "."},
				[]string{"-", "-", "-", "-", "5", ".", "#", "#", "#", "#", "#", ".", ".", ".", "."},
				[]string{"-", "-", "2", "4", "1", "#", "#", ".", "#", "#", "#", "#", ".", ".", "#"},
				[]string{"-", "-", "-", "1", "7", "#", ".", ".", "#", "#", "#", "#", "#", "#", "#"},
				[]string{"-", "-", "1", "5", "1", "#", ".", ".", "#", "#", "#", "#", "#", ".", "#"},
				[]string{"-", "-", "-", "-", "7", "#", "#", "#", "#", "#", "#", "#", ".", ".", "."},
				[]string{"-", "-", "-", "-", "5", ".", "#", "#", "#", "#", "#", ".", ".", ".", "."},
				[]string{"-", "-", "1", "1", "1", "#", ".", "#", ".", "#", ".", ".", ".", ".", "."},
				[]string{"-", "-", "-", "-", "7", "#", "#", "#", "#", "#", "#", "#", ".", ".", "."},
			},
			true,
		},
	}

	for i, test := range tests {
		actual := correctNonogram(test.size, test.n)
		if actual != test.x {
			t.Errorf("%d - Got %t; Wanted %t", i+1, actual, test.x)
		}
	}
}
