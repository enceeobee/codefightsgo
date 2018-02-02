package springofintegration

import (
	"testing"
)

func TestStringsCrossover(t *testing.T) {
	tests := []struct {
		i []string
		r string
		x int
	}{
		{
			[]string{
				"abc",
				"aaa",
				"aba",
				"bab"},
			"bbb",
			2,
		},
		{
			[]string{
				"aacccc",
				"bbcccc"},
			"abdddd",
			0,
		},
		{
			[]string{
				"a",
				"b",
				"c",
				"d",
				"e"},
			"c",
			4,
		},
		{
			[]string{
				"aa",
				"ab",
				"ba"},
			"bb",
			1,
		},
		{
			[]string{
				"aaa",
				"aaa"},
			"aaaf",
			1,
		},
	}

	for _, test := range tests {
		actual := stringsCrossover(test.i, test.r)
		if actual != test.x {
			t.Errorf("BONK. stringsCrossover(%q, %s) = %d, wanted %d", test.i, test.r, actual, test.x)
		}
	}
}
