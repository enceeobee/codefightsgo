package sortingoutpost

import (
	"reflect"
	"sort"
)

// Find the order for the sorted columns, then compare the orders
func rowsRearranging(matrix [][]int) bool {
	orders := [][]int{}
	cols := [][]int{}
	sortedCols := [][]int{}

	for c := 0; c < len(matrix[0]); c++ {
		col := []int{}
		sortedCol := []int{}

		for r := 0; r < len(matrix); r++ {
			col = append(col, matrix[r][c])
			sortedCol = append(sortedCol, matrix[r][c])
		}

		sort.Ints(sortedCol)

		cols = append(cols, col)
		sortedCols = append(sortedCols, sortedCol)
	}

	// Ensure there aren't dupes
	for _, sc := range sortedCols {
		for i, val := range sc {
			if i > 0 && sc[i-1] == val {
				// We have a duplicate value, return false
				return false
			}
		}
	}

	for i, sc := range sortedCols {
		order := []int{}

		for _, val := range sc {
			order = append(order, indexOf(cols[i], val))
		}

		orders = append(orders, order)
	}

	for i := 1; i < len(orders); i++ {
		if !reflect.DeepEqual(orders[0], orders[i]) {
			return false
		}
	}

	return true
}

func indexOf(s []int, v int) int {
	for i, n := range s {
		if n == v {
			return i
		}
	}

	return -1
}
