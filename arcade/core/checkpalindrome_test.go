package codefightsgo

import (
	"testing"
)

func TestCheckPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"hello", false},
		{"racecar", true},
		{"tacocat", true},
		{"anna", true},
	}
	for _, c := range cases {
		got := CheckPalindrome(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}
