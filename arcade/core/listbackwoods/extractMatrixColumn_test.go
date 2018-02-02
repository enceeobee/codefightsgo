package listbackwoods

import "testing"
import "reflect"

func TestExtractColumn(t *testing.T) {
	tests := []struct {
		matrix [][]int
		col    int
		x      []int
	}{
		{
			[][]int{
				[]int{1, 1, 1, 2},
				[]int{0, 5, 0, 4},
				[]int{2, 1, 3, 6},
			},
			2,
			[]int{1, 0, 3},
		},
	}

	for _, test := range tests {
		actual := extractMatrixColumn(test.matrix, test.col)

		if !reflect.DeepEqual(test.x, actual) {
			t.Errorf("nope")
		}
	}

}
