package codefightsgo

import "math"

// AdjacentElementsProduct finds the greatest product of any two adjacent values in inputArray.
func AdjacentElementsProduct(inputArray []int) int {
	greatestProduct := float64(inputArray[0] * inputArray[1])
	for i := 1; i < len(inputArray)-1; i++ {
		greatestProduct = math.Max(greatestProduct, float64(inputArray[i]*inputArray[i+1]))
	}
	return int(greatestProduct)
}
