package waterfallofintegration

func polygonPerimeter(matrix [][]bool) int {
	p := 0
	numCols := len(matrix[0])
	numRows := len(matrix)

	for r, row := range matrix {
		for c, cell := range row {
			if cell {
				// Check above
				if r == 0 || !matrix[r-1][c] {
					p++
				}
				// Check right
				if c == numCols-1 || !matrix[r][c+1] {
					p++
				}
				// Check below
				if r == numRows-1 || !matrix[r+1][c] {
					p++
				}
				// Check left
				if c == 0 || !matrix[r][c-1] {
					p++
				}
			}
		}
	}
	return p
}
