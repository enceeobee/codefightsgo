package regularhell

import "testing"

func TestReplaceAllDigitsRegExp(t *testing.T) {
	tests := []struct {
		input string
		x     string
	}{
		{
			"There are 12 points",
			"There are ## points",
		},
		{
			"012ss210",
			"###ss###",
		},
		{
			" _Aas 23",
			" _Aas ##",
		},
		{
			"no digits here",
			"no digits here",
		},
		{
			" aa_0239mehce3d",
			" aa_####mehce#d",
		},
		{
			"v z gv4zyx eu mu ",
			"v z gv#zyx eu mu ",
		},
	}

	var actual string
	for _, test := range tests {
		actual = replaceAllDigitsRegExp(test.input)

		if actual != test.x {
			t.Errorf("replaceAllDigitsRegExp(%s) = %s; wanted %s", test.input, actual, test.x)
		}
	}
}
