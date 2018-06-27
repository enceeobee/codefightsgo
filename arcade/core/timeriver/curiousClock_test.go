package timeriver

import "testing"

func TestCuriousClock(t *testing.T) {
	var actual string
	tests := []struct {
		some    string
		leaving string
		x       string
	}{
		{"2016-08-26 22:40", "2016-08-29 10:00", "2016-08-24 11:20"},
		{"2016-08-26 22:40", "2016-08-26 22:41", "2016-08-26 22:39"},
		{"2015-01-14 09:12", "2015-11-04 17:36", "2014-03-26 00:48"},
		{"2016-02-28 12:21", "2016-03-01 22:21", "2016-02-26 02:21"},
		{"1995-05-30 13:48", "2016-04-22 05:32", "1974-07-06 22:04"},
		{"1992-09-29 00:00", "1998-12-04 23:59", "1986-07-25 00:01"},
	}

	for _, test := range tests {
		actual = curiousClock(test.some, test.leaving)

		if actual != test.x {
			t.Errorf("curiousClock(%q, %q) = %q; expected %q", test.some, test.leaving, actual, test.x)
		}
	}
}
