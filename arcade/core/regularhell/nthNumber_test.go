package regularhell

import "testing"

func TestNthNumber(t *testing.T) {
	tests := []struct {
		s string
		n int
		x string
	}{
		{"8one 003number 201numbers li-000233le number444", 4, "233"},
		{"some023020 num ber 033 02103 32 meh peh beh 4328 ", 5, "4328"},
		{"007 is an awesome agent", 1, "7"},
		{"Everyone hates error 404", 1, "404"},
		{"LaS003920tP3rEt4t04Yte0023s3t", 4, "4"},
		{"=_aaYiM*}&0077|D))ztIV00012432748730156644204805614898523099655216oio0854102350044141_|YDL0584511290939644184700867021026771007612398051168360441323oIc:G*0204864749576405915~wqgN0037594999439741539584332{F&wjxiLy-mE", 4, "584511290939644184700867021026771007612398051168360441323"},
	}

	var actual string

	for _, test := range tests {
		actual = nthNumber(test.s, test.n)

		if actual != test.x {
			t.Errorf("nthNumber(%q, %d) = %q; expected %q", test.s, test.n, actual, test.x)
		}
	}
}
