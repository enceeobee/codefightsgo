package timeriver

import "testing"

func TestMissedClasses(t *testing.T) {
	var actual int
	tests := []struct {
		year          int
		daysOfTheWeek []int
		holidays      []string
		x             int
	}{
		{
			2015,
			[]int{2, 3},
			[]string{"11-04",
				"02-22",
				"02-23",
				"03-07",
				"03-08",
				"05-09"},
			3,
		},
		{
			1900,
			[]int{},
			[]string{},
			0,
		},
		{
			2100,
			[]int{1, 4, 7},
			[]string{"10-28",
				"05-03",
				"10-02",
				"05-07",
				"05-25",
				"09-04",
				"10-30",
				"03-03",
				"09-02",
				"11-08"},
			4,
		},
		{
			1956,
			[]int{1, 4, 6, 7},
			[]string{"03-17",
				"04-03",
				"03-08",
				"09-18",
				"05-28",
				"02-14",
				"10-20",
				"01-02",
				"01-22",
				"10-04",
				"02-02",
				"10-07",
				"09-30",
				"05-20",
				"05-23",
				"09-22",
				"01-12",
				"05-02",
				"10-21",
				"03-20"},
			13,
		},
	}

	for i, test := range tests {
		actual = missedClasses(test.year, test.daysOfTheWeek, test.holidays)

		if actual != test.x {
			t.Errorf("Test %d - Got %d; Wanted %d", i+1, actual, test.x)
		}
	}
}
