package timeriver

import "testing"

func TestHoliday(t *testing.T) {
	tests := []struct {
		x          int
		weekDay    string
		month      string
		yearNumber int
		expected   int
	}{
		{3, "Wednesday", "November", 2016, 16},
		{5, "Thursday", "November", 2016, -1},
		{1, "Thursday", "January", 2015, 1},
		{2, "Thursday", "January", 2015, 8},
		{3, "Thursday", "January", 2015, 15},
		{3, "Thursday", "January", 2101, 20},
		{3, "Thursday", "January", 2401, 18},
		{2, "Thursday", "December", 2500, 9},
		{5, "Monday", "February", 2016, 29},
	}

	var actual int
	for i, test := range tests {
		actual = holiday(test.x, test.weekDay, test.month, test.yearNumber)

		if actual != test.expected {
			t.Errorf("Test %d = %d; expected %d", i+1, actual, test.expected)
		}
	}
}
