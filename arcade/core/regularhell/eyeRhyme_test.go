package regularhell

import "testing"

func TestEyeRhyme(t *testing.T) {
	tests := []struct {
		p string
		x bool
	}{
		{"cough\tbough", true},
		{"CodeFig!ht\tWith all your might", false},
		{"quod erat demonstrandum\tand that, ladies and gentlemen, is the end of my memorandum", true},
		{"yup\tyes", false},
		{"hey\they", true},
		{"E = MC<sup>2</sup>\twhich in turn equals sup", false},
		{"Isn't it correct?!\tIt definitely is! Does it mean that we've just obtained a full set?!", true},
		{"Nothing can go wrong here :)\tHehehehe:)", false},
		{"!1?/\tsldkjfl3 sldjf 1?/", true},
		{"simple\tpimple", true},
		// {"tray\tcraycray", true},
	}

	var actual bool

	for _, test := range tests {
		actual = eyeRhyme(test.p)

		if actual != test.x {
			t.Errorf("eyeRhyme(%s) = %t; expected %t", test.p, actual, test.x)
		}
	}
}
