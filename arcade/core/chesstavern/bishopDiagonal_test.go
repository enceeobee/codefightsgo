package chesstavern

import (
	"reflect"
	"testing"
)

func TestBishopDiagonal(t *testing.T) {
	// actual := []string{}
	tests := []struct {
		b1 string
		b2 string
		x  []string
	}{
		{
			"d7",
			"f5",
			[]string{"c8", "h3"},
		},
		{
			"d8",
			"b5",
			[]string{"b5", "d8"},
		},
		{
			"a1",
			"h8",
			[]string{"a1", "h8"},
		},
		{
			"g3",
			"e1",
			[]string{"e1", "h4"},
		},
		{
			"b4",
			"e7",
			[]string{"a3", "f8"},
		},
	}

	for i, test := range tests {
		actual := bishopDiagonal(test.b1, test.b2)

		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("%d - Got %v; Wanted %v", i+1, actual, test.x)
		}
	}
}
