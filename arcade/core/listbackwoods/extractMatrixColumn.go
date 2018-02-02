package listbackwoods

func extractMatrixColumn(matrix [][]int, column int) []int {
	col := []int{}
	for _, row := range matrix {
		col = append(col, row[column])
	}
	return col
}
