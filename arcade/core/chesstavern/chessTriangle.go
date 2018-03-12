package chesstavern

func chessTriangle(n int, m int) int {
	if (n < 2 || m < 2) || (n*m < 6) {
		return 0
	}

	/*
		Triangles:

		* Need at least a 2x3 grid
		* Knight is always in corner
		* We can fit 8 triangles in a 2x3 grid
		* Can stretch with the bishop:
			* The following:
				* k x x
				* x b r
			* can also be:
				* k x x
				* x x r
				* x x b
			* SO, when bishop takes knight, we get an extra coupla triangles


	*/
	return 3
}
