package chesstavern

import (
	"reflect"
	"testing"
)

func TestChessBishopDream(t *testing.T) {
	tests := []struct {
		boardSize []int
		initPos   []int
		initDir   []int
		k         int
		x         []int
	}{
		{
			[]int{3, 7},
			[]int{1, 2},
			[]int{-1, 1},
			13,
			[]int{0, 1},
		},
		{
			[]int{1, 2},
			[]int{0, 0},
			[]int{1, 1},
			6,
			[]int{0, 1},
		},
		{
			[]int{2, 2},
			[]int{1, 0},
			[]int{1, 1},
			12,
			[]int{1, 0},
		},
		{
			[]int{1, 1},
			[]int{0, 0},
			[]int{1, -1},
			1000000000,
			[]int{0, 0},
		},
		{
			[]int{2, 3},
			[]int{1, 2},
			[]int{-1, -1},
			41,
			[]int{0, 2},
		},
		{
			[]int{17, 19},
			[]int{14, 8},
			[]int{1, -1},
			239239,
			[]int{4, 17},
		},
		{
			[]int{17, 19},
			[]int{16, 18},
			[]int{1, 1},
			239239239,
			[]int{10, 2},
		},
	}

	for _, test := range tests {
		actual := chessBishopDream(test.boardSize, test.initPos, test.initDir, test.k)

		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("chessBishopDream(%v, %v, %v, %d) wanted %v; got %v", test.boardSize, test.initPos, test.initDir, test.k, test.x, actual)
		}
	}
}
