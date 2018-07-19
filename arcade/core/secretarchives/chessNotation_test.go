package secretarchives

import "testing"

func TestChessNotation(t *testing.T) {
	tests := []struct {
		notation string
		x        string
	}{
		{"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR", "RP4pr/NP4pn/BP4pb/QP4pq/K2P2pk/BP4pb/NP4pn/RP4pr"},
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", "RP4pr/NP4pn/BP4pb/QP4pq/KP4pk/BP4pb/NP4pn/RP4pr"},
		{"2kr3r/pp1nbppp/3p1n2/q1pPp1B1/4P1b1/2N2N2/PPP1BPPP/R2Q2RK", "RP2q1p1/1P4p1/1PN1p2k/Q3Ppnr/1B1Pp1b1/1PN2np1/RP1bB1p1/KP4pr"},
		{"8/8/8/8/8/8/8/8", "8/8/8/8/8/8/8/8"},
		{"8/n2Q3R/8/8/2k3N1/2b1r2q/2B5/8", "6n1/8/1Bbk4/6Q1/2r5/8/3N4/2q3R1"},
	}

	var actual string
	for _, test := range tests {
		actual = chessNotation(test.notation)
		if actual != test.x {
			t.Errorf("\nGot\n%q\nexpected\n%q", actual, test.x)
		}
	}
}
