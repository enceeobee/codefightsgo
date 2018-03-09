package chesstavern

import "testing"

func TestChessKnightMoves(t *testing.T) {
	var actual int
	tests := []struct {
		cell string
		x    int
	}{
		{"a1", 2},
		{"c2", 6},
		{"h8", 2},
	}

	for i, test := range tests {
		actual = chessKnightMoves(test.cell)

		if actual != test.x {
			t.Errorf("%d - Got %d; Wanted %d", i+1, actual, test.x)
		}
	}
}
