package waterfallofintegration

import (
	"reflect"
	"testing"
)

func TestGravitation(t *testing.T) {
	tests := []struct {
		rows []string
		x    []int
	}{
		// 1
		{
			[]string{
				"#..##",
				".##.#",
				".#.##",
				"....."},
			[]int{1, 4},
		},
		// 2
		{
			[]string{
				"#..##",
				".##.#",
				".#.##",
				"..##."},
			[]int{1, 2, 3, 4},
		},
		// 3
		{
			[]string{
				"##",
				"..",
				"..",
				".."},
			[]int{0, 1},
		},
		// 4
		{
			[]string{
				"#..#.",
				".##..",
				".#...",
				".#..."},
			[]int{1, 4},
		},
		// 5
		{
			[]string{
				"#....",
				".#..#",
				"..#..",
				"...#."},
			[]int{3},
		},

		// CUSTOM, BIATCH
		{
			[]string{
				"#.###",
				".#...",
				"#...."},
			[]int{0, 1},
		},
	}

	for i, test := range tests {
		actual := gravitation(test.rows)
		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("#%d - Wanted %v; Got %v", i+1, test.x, actual)
		}
	}
}
