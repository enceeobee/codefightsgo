package chesstavern

import "testing"

func TestWhoseTurn(t *testing.T) {
	tests := []struct {
		p string
		x bool
	}{
		{"b1;g1;b8;g8", true},
		{"c3;g1;b8;g8", false},
		{"g1;g2;g3;g4", true},
		{"f8;h1;f3;c2", false},
	}

	var actual bool
	for _, test := range tests {
		actual = whoseTurn(test.p)
		if actual != test.x {
			t.Errorf("whoseTurn(%s) = %t; wanted %t", test.p, actual, test.x)
		}
	}
}
