package springofintegration

import "testing"

func TestCombs(t *testing.T) {
	tests := []struct {
		comb1 string
		comb2 string
		ex    int
	}{
		{"*..*", "*.*", 5},
		{"*...*", "*.*", 5}, // = **.**
		{"*.*", "*...*", 5},
		{"*..*.*", "*.***", 9},
		{"*.*", "*.*", 4},
		{"*.**", "*.*", 5},
	}

	for _, test := range tests {
		actual := combs(test.comb1, test.comb2)

		if actual != test.ex {
			t.Errorf("combs bonked - got %d, wanted %d", actual, test.ex)
		}
	}
}

// func TestGetLen(t *testing.T) {
// 	combs = [][]byte{
// 		[]byte{byte('*'), byte('.'), byte('*')},
// 		[]byte{byte('*'), byte('.'), byte('*')},
// 	}
// }
