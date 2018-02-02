package codefightsgo

import "testing"

func TestReverseParentheses(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"a(bc)de", "acbde"},
		{"a(bcdefghijkl(mno)p)q", "apmnolkjihgfedcbq"},
		{"co(de(fight)s)", "cosfighted"},
		{"Code(Cha(lle)nge)", "CodeegnlleahC"},
		{"Where are the parentheses", "Where are the parentheses"},
		{"abc(cba)ab(bac)c", "abcabcabcabc"},
		{"The ((quick (brown) (fox) jumps over the lazy) dog)", "The god quick nworb xof jumps over the lazy"},
	}

	for _, test := range tests {
		actual := ReverseParentheses(test.in)
		if actual != test.expected {
			t.Errorf("For %q; expected %q, got %q", test.in, test.expected, actual)
		}
	}
}
