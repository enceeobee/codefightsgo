package listbackwoods

func crossingSum(matrix [][]int, a int, b int) int {
	var rowSum, colSum int

	for _, n := range matrix[a] {
		rowSum += n
	}

	for i := 0; i < len(matrix); i++ {
		if i == a {
			continue
		}
		colSum += matrix[i][b]
	}

	return rowSum + colSum
}
