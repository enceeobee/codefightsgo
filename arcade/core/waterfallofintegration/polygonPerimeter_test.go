package waterfallofintegration

import (
	"reflect"
	"testing"
)

func TestPolygonPerimeter(t *testing.T) {
	tests := []struct {
		matrix [][]bool
		x      int
	}{
		{
			[][]bool{
				[]bool{false, true, true},
				[]bool{true, true, false},
				[]bool{true, false, false},
			},
			12,
		},

		{
			[][]bool{
				[]bool{true, true, true},
				[]bool{true, false, true},
				[]bool{true, true, true},
			},
			16,
		},

		{
			[][]bool{
				[]bool{true, true, true, true, true},
				[]bool{true, true, true, true, true},
			},
			14,
		},
	}

	for i, test := range tests {
		actual := polygonPerimeter(test.matrix)

		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("%d failed", i+1)
		}
	}
}
