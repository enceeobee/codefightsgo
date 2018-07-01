package timeriver

import "testing"

func TestRegularMonths(t *testing.T) {
	var actual string
	tests := []struct {
		curMonth string
		x        string
	}{
		{"02-2016", "08-2016"},
		{"05-2027", "11-2027"},
		{"09-2099", "02-2100"},
		{"01-1970", "06-1970"},
		{"07-2024", "09-2025"},
	}

	for i, test := range tests {
		actual = regularMonths(test.curMonth)

		if actual != test.x {
			t.Errorf("Test %d - Got %s; Wanted %s", i+1, actual, test.x)
		}
	}
}
