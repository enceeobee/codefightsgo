package codefightsgo

import (
	"reflect"
	"testing"
)

func TestAllLongestStrings(t *testing.T) {
	tests := []struct {
		in       []string
		expected []string
	}{
		{
			[]string{"aba", "aa"},
			[]string{"aba"},
		},
	}

	for _, test := range tests {
		actual := AllLongestStrings(test.in)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("AllLongestStrings(%q) == %q, want %q", test.in, actual, test.expected)
		}
	}
}
