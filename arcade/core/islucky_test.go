package codefightsgo

import "testing"

func TestIsLucky(t *testing.T) {
	tests := []struct {
		in       int
		expected bool
	}{
		{1230, true},
		{239017, false},
		{134008, true},
		{10, false},
		{11, true},
		{1010, true},
		{261534, false},
		{100000, false},
		{999999, true},
		{123321, true},
	}

	for _, test := range tests {
		actual := IsLucky(test.in)
		if actual != test.expected {
			t.Errorf("For %d, expected %t, got %t", test.in, test.expected, actual)
		}
	}
}
