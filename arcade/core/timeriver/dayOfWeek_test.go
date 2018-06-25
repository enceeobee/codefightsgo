package timeriver

import "testing"

func TestDayOfWeek(t *testing.T) {
	tests := []struct {
		b string
		x int
	}{
		{"02-01-2016", 5},
		{"01-01-1900", 6},
		{"02-29-2016", 28},
		{"01-16-2027", 11},
		{"10-12-2000", 6},
		{"02-29-2072", 40},
	}

	for _, test := range tests {
		actual := dayOfWeek(test.b)
		if actual != test.x {
			t.Errorf("dayOfWeek(%s) = %d; expected %d", test.b, actual, test.x)
		}
	}
}
