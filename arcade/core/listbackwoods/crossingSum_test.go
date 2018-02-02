package listbackwoods

import "testing"

func TestCrossingSum(t *testing.T) {
	tests := []struct {
		matrix [][]int
		a      int
		b      int
		x      int
	}{
		{
			[][]int{[]int{1, 1, 1, 1}, []int{2, 2, 2, 2}, []int{3, 3, 3, 3}},
			1,
			3,
			12,
		},
	}

	for _, test := range tests {
		actual := crossingSum(test.matrix, test.a, test.b)
		if actual != test.x {
			t.Errorf("Wanted %d; got %d", test.x, actual)
		}
		// if !reflect.DeepEqual(test.x, actual) {
		// 	t.Errorf("bahh")
		// }
	}
}
