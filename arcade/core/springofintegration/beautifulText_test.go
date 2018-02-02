package springofintegration

import (
	"testing"
)

func TestBeautifulText(t *testing.T) {
	tests := []struct {
		i string
		l int
		r int
		x bool
	}{
		{"Look at this example of a correct text", 5, 15, true},
		{"Look at this example of a correct text", 12, 13, true},
		{"abc def ghi", 4, 10, false},
		{"a a a a a a a a", 1, 10, true},
		{"ab cd fg xyz", 1, 5, false},
		{"aa aa aaaaa aaaaa aaaaa", 6, 11, true},
		{"aa aa aaaaa aaaaa aaaaa", 5, 10, true},
		{"aa aa aaaaa aaaaa aaaaa", 6, 10, false},
		{"aaa aaaaaaa", 3, 10, false},
		{"dht skq dkg", 4, 10, false},
		{"skspv iakqh ygzwz ntkmi xqhpj", 3, 7, true},
		{"skspv iakqh ygzwz ntkmi xqhpj", 8, 9, false},
	}

	var actual bool

	for _, test := range tests {
		actual = beautifulText(test.i, test.l, test.r)

		if actual != test.x {
			t.Errorf("beautifulText(%q, %d, %d) = %t; wanted %t", test.i, test.l, test.r, actual, test.x)
		}
	}
}
