package listbackwoods

import "testing"
import "reflect"

func TestSwapDiagonals(t *testing.T) {
	tests := []struct {
		matrix [][]int
		x      [][]int
	}{
		{
			[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
			[][]int{[]int{3, 2, 1}, []int{4, 5, 6}, []int{9, 8, 7}},
		},
	}

	for _, test := range tests {
		actual := swapDiagonals(test.matrix)
		if !reflect.DeepEqual(test.x, actual) {
			t.Errorf("bahh")
		}
	}
}
