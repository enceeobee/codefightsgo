package listbackwoods

import "testing"
import "reflect"

func TestReverseOnDiagonals(t *testing.T) {
	tests := []struct {
		matrix [][]int
		x      [][]int
	}{
		{
			[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
			[][]int{[]int{9, 2, 7}, []int{4, 5, 6}, []int{3, 8, 1}},
		},
	}

	for _, test := range tests {
		actual := reverseOnDiagonals(test.matrix)
		if !reflect.DeepEqual(test.x, actual) {
			t.Errorf("bahh")
		}
	}
}
