package listbackwoods

import "testing"
import "reflect"

func TestDrawRectangle(t *testing.T) {
	tests := []struct {
		canvas    [][]string
		rectangle []int
		x         [][]string
	}{
		{
			[][]string{
				[]string{"a", "a", "a", "a", "a", "a", "a", "a"},
				[]string{"a", "a", "a", "a", "a", "a", "a", "a"},
				[]string{"a", "a", "a", "a", "a", "a", "a", "a"},
				[]string{"b", "b", "b", "b", "b", "b", "b", "b"},
				[]string{"b", "b", "b", "b", "b", "b", "b", "b"},
			},
			[]int{1, 1, 4, 3},
			[][]string{
				[]string{"a", "a", "a", "a", "a", "a", "a", "a"},
				[]string{"a", "*", "-", "-", "*", "a", "a", "a"},
				[]string{"a", "|", "a", "a", "|", "a", "a", "a"},
				[]string{"b", "*", "-", "-", "*", "b", "b", "b"},
				[]string{"b", "b", "b", "b", "b", "b", "b", "b"},
			},
		},
	}

	for _, test := range tests {
		actual := drawRectangle(test.canvas, test.rectangle)
		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("arrrjjjjj")
		}
	}
}
