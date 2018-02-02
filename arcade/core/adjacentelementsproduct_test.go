package codefightsgo

import (
	"testing"
)

func TestAdjacentElementsProduct(t *testing.T) {
	cases := []struct {
		in       []int
		expected int
	}{
		{[]int{3, 6, -2, -5, 7, 3}, 21},
		{[]int{-1, -2}, 2},
		{[]int{5, 1, 2, 3, 1, 4}, 6},
		{[]int{1, 2, 3, 0}, 6},
		{[]int{9, 5, 10, 2, 24, -1, -48}, 50},
		{[]int{5, 6, -4, 2, 3, 2, -23}, 30},
		{[]int{4, 1, 2, 3, 1, 5}, 6},
		{[]int{-23, 4, -3, 8, -12}, -12},
	}
	for _, c := range cases {
		actual := AdjacentElementsProduct(c.in)
		if actual != c.expected {
			t.Errorf("Reverse(%d) == %d, expected %d", c.in, actual, c.expected)
		}
	}
}
