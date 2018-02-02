package listbackwoods

func reverseOnDiagonals(matrix [][]int) [][]int {
	var temp int
	lenMatrix := len(matrix)

	for i := 0; i < lenMatrix/2; i++ {
		opp := lenMatrix - 1 - i
		temp = matrix[i][i]

		matrix[i][i] = matrix[opp][opp]
		matrix[opp][opp] = temp

		temp = matrix[i][opp]
		matrix[i][opp] = matrix[opp][i]
		matrix[opp][i] = temp

	}

	return matrix
}
