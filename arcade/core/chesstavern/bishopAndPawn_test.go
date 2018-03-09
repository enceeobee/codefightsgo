package chesstavern

import "testing"

func TestBishopAndPawn(t *testing.T) {
	var actual bool
	tests := []struct {
		b string
		p string
		x bool
	}{
		{"a1", "c3", true},
		{"h1", "h3", false},
		{"a5", "c3", true},
		{"g1", "f3", false},
		{"e7", "d6", true},
		{"e7", "a3", true},
		{"e3", "a7", true},
		{"a1", "h8", true},
		{"a1", "h7", false},
		{"h1", "a8", true},
	}

	for i, test := range tests {
		actual = bishopAndPawn(test.b, test.p)
		if actual != test.x {
			t.Errorf("%d - Wanted %t; Got %t", i+1, test.x, actual)
		}
	}
}
