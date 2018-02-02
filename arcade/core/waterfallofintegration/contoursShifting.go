package waterfallofintegration

import (
	"math"
)

type contour struct {
	index   int
	numRows int
	numCols int
	values  []int
}

func contoursShifting(matrix [][]int) [][]int {
	matrixcopy := make([][]int, len(matrix))

	for i := range matrix {
		matrixcopy[i] = make([]int, len(matrix[i]))
		copy(matrixcopy[i], matrix[i])
	}

	contours := []contour{}
	numRows := len(matrix)
	numCols := len(matrix[0])
	min := int(math.Min(float64(numRows), float64(numCols)))
	index := 0
	var tempcontour contour

	for i := min; i > 0; i -= 2 {
		tempcontour.index = index
		tempcontour.numRows = numRows - (index * 2)
		tempcontour.numCols = numCols - (index * 2)
		tempcontour.values = []int{}

		// Top row
		for tri := 0; tri < tempcontour.numCols; tri++ {
			tempcontour.values = append(tempcontour.values, matrix[index][index+tri])
		}

		if tempcontour.numRows > 1 {
			// Right col
			for rci := 0; rci < tempcontour.numRows-1; rci++ {
				tempcontour.values = append(tempcontour.values, matrix[rci+index+1][tempcontour.numCols+index-1])
			}

			if tempcontour.numCols > 1 {
				// Bottom row
				for bri := 0; bri < tempcontour.numCols-1; bri++ {
					tempcontour.values = append(tempcontour.values, matrix[tempcontour.numRows-1+index][tempcontour.numCols-2-bri+index])
				}

				// Left col
				for lci := 0; lci < tempcontour.numRows-2; lci++ {
					tempcontour.values = append(tempcontour.values, matrix[tempcontour.numRows-2-lci+index][index])
				}
			}
		}

		// Shift contour
		if index%2 == 0 {
			tempcontour.values = shiftClockwise(tempcontour.values)
		} else {
			tempcontour.values = shiftCounterClockwise(tempcontour.values)
		}

		contours = append(contours, tempcontour)
		index++
	}

	// Put the matrix back together again
	for index, contr := range contours {
		vali := 0

		// Top row
		for tri := index; tri < contr.numCols+index; tri++ {
			matrixcopy[index][tri] = contr.values[vali]
			vali++
		}

		if contr.numRows > 1 {
			// Right col
			for rci := 0; rci < contr.numRows-1; rci++ {
				matrixcopy[rci+index+1][contr.numCols+index-1] = contr.values[vali]
				vali++
			}

			if contr.numCols > 1 {
				// Bottom row
				for bri := 0; bri < contr.numCols-1; bri++ {
					matrixcopy[contr.numRows-1+index][contr.numCols-2-bri+index] = contr.values[vali]
					vali++
				}

				// Left col
				for lci := 0; lci < contr.numRows-2; lci++ {
					matrixcopy[contr.numRows-2-lci+index][index] = contr.values[vali]
					vali++
				}
			}
		}
	}

	return matrixcopy
}

func shiftClockwise(values []int) []int {
	return append([]int{values[len(values)-1]}, values[0:len(values)-1]...)
}

func shiftCounterClockwise(values []int) []int {
	return append(append([]int{values[1]}, values[2:len(values)]...), values[0])
}
