package regularhell

import "testing"

func TestSwapAdjacentWords(t *testing.T) {
	tests := []struct {
		s string
		x string
	}{
		{
			"CodeFight On",
			"On CodeFight",
		},
		{
			"How are you today guys",
			"are How today you guys",
		},
		{
			"IAmALongStringWithNoWhiteSpaceCharacters",
			"IAmALongStringWithNoWhiteSpaceCharacters",
		},
		{
			"A b C D e F g h I j",
			"b A D C F e h g j I",
		},
		{
			"",
			"",
		},
	}

	var actual string
	for _, test := range tests {
		actual = swapAdjacentWords(test.s)

		if actual != test.x {
			t.Errorf("swapAdjacentWords(%s) = %s; expected %s", test.s, actual, test.x)
		}
	}
}
