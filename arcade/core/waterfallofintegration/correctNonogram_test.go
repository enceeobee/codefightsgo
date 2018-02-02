package waterfallofintegration

import "testing"

func TestCorrectNonogram(t *testing.T) {
	tests := []struct {
		size int
		n    [][]string
		x    bool
	}{
		{
			5,
			[][]string{},
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
