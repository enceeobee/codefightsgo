package regularhell

import (
	"testing"
)

func TestBugsAndBugfixes(t *testing.T) {
	var actual int
	tests := []struct {
		r string
		x int
	}{
		{"Roll d6-3 and 4d4+3 to pick a weapon, and finish the boss with 3d7!", 43},
		{"d6-almost 01d4+2 12d01-3 5d5-00 a valid formula", 46},
		{"meh4d2-3D3", 5},
		{"ad3+4, 44b-6, 5daa", 7},
		{"4d6-L1d20-10 did4n't expect that", 38},
		{"nothing here", 0},
	}

	for _, test := range tests {
		actual = bugsAndBugfixes(test.r)
		if actual != test.x {
			t.Errorf("bugsAndBugfixes(%s) = %d; expected %d\n", test.r, actual, test.x)
		}
	}
}
