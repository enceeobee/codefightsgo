package listbackwoods

func drawRectangle(canvas [][]string, rectangle []int) [][]string {
	x1 := rectangle[0]
	y1 := rectangle[1]
	x2 := rectangle[2]
	y2 := rectangle[3]

	// Set corners
	canvas[y1][x1] = "*"
	canvas[y1][x2] = "*"
	canvas[y2][x1] = "*"
	canvas[y2][x2] = "*"

	// Draw top and bottom
	for i := x1 + 1; i < x2; i++ {
		canvas[y1][i] = "-"
		canvas[y2][i] = "-"
	}

	// Draw left and right
	for i := y1 + 1; i < y2; i++ {
		canvas[i][x1] = "|"
		canvas[i][x2] = "|"
	}

	return canvas
}
