package springofintegration

import "testing"

func TestCyclicStrings(t *testing.T) {
	var actual int
	tests := []struct {
		s string
		x int
	}{
		{"cabca", 3},
		{"aba", 2},
		{"ccccccccccc", 1},
		{"baa", 3},
		{"abaa", 3},
		{"aab", 3},
		{"aabaa", 3},
		{"aabaab", 3},
	}

	for _, test := range tests {
		actual = cyclicString(test.s)

		if actual != test.x {
			t.Errorf("cyclicString(%s) = %d; wanted %d", test.s, actual, test.x)
		}
	}
}
