package codefightsgo

import (
	"testing"
)

func TestCommonCharacterCount(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected int
	}{
		{"aabcc", "adcaa", 3},
		{"zzzz", "zzzzzzz", 4},
		{"abca", "xyzbac", 3},
		{"a", "b", 0},
		{"a", "aaa", 1},
	}

	for _, test := range tests {
		actual := CommonCharacterCount(test.a, test.b)
		if actual != test.expected {
			t.Errorf("For %q, %q; expected %d, got %d", test.a, test.b, test.expected, actual)
		}
	}
}
