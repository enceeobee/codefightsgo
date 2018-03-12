package chesstavern

import "testing"

func TestChessTriangle(t *testing.T) {
	tests := []struct {
		n int
		m int
		x int
	}{
		{2, 3, 8},
		{1, 30, 0},
		{3, 3, 48},
		{2, 2, 0},
		{5, 2, 40},
	}

	var actual int
	for _, test := range tests {
		actual = chessTriangle(test.n, test.m)
		if actual != test.x {
			t.Errorf("chessTriangle(%d, %d) = %d; wanted %d", test.n, test.m, actual, test.x)
		}
	}
}
