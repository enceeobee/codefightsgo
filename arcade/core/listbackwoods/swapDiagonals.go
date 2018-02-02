package listbackwoods

import "fmt"

func swapDiagonals(matrix [][]int) [][]int {
	var temp int

	/*
		diags for a 3x3:
		0,0; 0,2
		1,1; 1,1
		2,2; 2,0
	*/

	for i := 0; i < len(matrix)/2; i++ {
		opp := len(matrix) - 1 - i

		temp = matrix[i][i]
		// 0,0 becomes 0,2
		matrix[i][i] = matrix[i][opp]
		// 0,2 becomes 0,0
		matrix[i][opp] = temp

		// set right
		temp = matrix[opp][opp]
		// 2,2 becomes 2,0
		matrix[opp][opp] = matrix[opp][i]
		//2,0 becomes 2,2
		matrix[opp][i] = temp

	}

	for _, row := range matrix {
		fmt.Println(row)
	}

	return matrix
}
