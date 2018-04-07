package timeriver

import "testing"

func TestValidTime(t *testing.T) {
	var actual bool
	tests := []struct {
		t string
		x bool
	}{
		{"13:58", true},
		{"25:51", false},
		{"02:76", false},
		{"24:00", false},
	}

	for _, test := range tests {
		actual = validTime(test.t)

		if actual != test.x {
			t.Errorf("validTime(%s) = %t; wanted %t", test.t, actual, test.x)
		}
	}
}
