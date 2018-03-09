package chesstavern

func chessKnightMoves(cell string) int {
	s := 0
	row := cell[1]
	col := cell[0]

	// Start at 1:00, and travel clockwise

	// Right side
	if col < 'h' && row < '7' {
		s++
	}
	if col < 'g' && row < '8' {
		s++
	}
	if col < 'g' && row > '1' {
		s++
	}
	if col < 'h' && row > '2' {
		s++
	}

	// Left side
	if col > 'a' && row > '2' {
		s++
	}
	if col > 'b' && row > '1' {
		s++
	}
	if col > 'b' && row < '8' {
		s++
	}
	if col > 'a' && row < '7' {
		s++
	}

	return s
}
