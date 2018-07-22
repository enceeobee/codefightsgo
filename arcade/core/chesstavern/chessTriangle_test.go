package chesstavern

import (
	"testing"
)

func TestChessTriangle(t *testing.T) {
	tests := []struct {
		h int
		w int
		x int
	}{
		{2, 3, 8},
		{1, 30, 0},
		{3, 3, 48},
		{2, 2, 0},
		{5, 2, 40},
		{3, 4, 112},
		{4, 3, 112},
		{7, 8, 1904},
		{18, 18, 16368},
		{40, 15, 31600},
		{25, 39, 54448},
		{40, 40, 92400},
	}

	var actual int
	for _, test := range tests {
		actual = chessTriangle(test.h, test.w)

		if actual != test.x {
			t.Errorf("chessTriangle(%d, %d) = %d; expected %d", test.h, test.w, actual, test.x)
		}
	}
}
