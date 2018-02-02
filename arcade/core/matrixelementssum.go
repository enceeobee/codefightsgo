package codefightsgo

func MatrixElementsSum(matrix [][]int) int {
	totalPrice := 0

	for col := 0; col < len(matrix[0]); col++ {
		for row := 0; row < len(matrix); row++ {
			totalPrice += matrix[row][col]
			if matrix[row][col] == 0 {
				break
			}
		}
	}

	return totalPrice
}
