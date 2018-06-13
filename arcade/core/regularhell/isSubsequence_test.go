package regularhell

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		t string
		s string
		x bool
	}{
		{"CodeFights", "CoFi", true},
		{"CodeFights", "cofi", false},
		{"somestring", "", true},
		{"penny", "longsenselessstringwithpennyinside", false},
		{"parliament", "partialmen", false},
		{"", "", true},
		{"", "hobo", false},
		{"he sd.f dsk e8.sd??l**23, 23,f.s++83+", "h  8.?*3+", true},
		{"abc", "d", false},
		{"abcd", "ad", true},
		{"longsenselessstringwithpennyinside", "penny", true},
		{"(test)", "test", true},
		{"((test))", "(test)", true},
	}

	var actual bool

	for _, test := range tests {
		actual = isSubsequence(test.t, test.s)

		if actual != test.x {
			t.Errorf("isSubsequence(%q, %q) = %t; expected %t", test.t, test.s, actual, test.x)
		}
	}
}
