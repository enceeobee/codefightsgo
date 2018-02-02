package listbackwoods

func volleyballPositions(formation [][]string, k int) [][]string {
	temp := []string{}
	players := []string{
		formation[0][1],
		formation[1][0],
		formation[1][2],
		formation[2][1],
		formation[3][0],
		formation[3][2],
	}

	/*
		0 3 0
		4 0 2
		0 6 0
		5 0 1
		= reg =
		0 4 0
		5 0 3
		0 1 0
		6 0 2

		0 5 0
		4 0 2
		0 3 0
		6 0 1
		=rev=
		0 2 0
		5 0 1
		0 6 0
		4 0 3


	*/
	// [Player5 Player4 Player2 Player3 Player6 Player1]
	// rev
	// [2 5 1 6 4 3]

	// NOTE - They rotate weirdly
	for i := 0; i < k%6; i++ {
		temp = []string{
			players[2],
			players[0],
			players[5],
			players[4],
			players[1],
			players[3],
		}

		copy(players, temp)
	}

	return [][]string{
		[]string{"empty", players[0], "empty"},
		[]string{players[1], "empty", players[2]},
		[]string{"empty", players[3], "empty"},
		[]string{players[4], "empty", players[5]},
	}
}
