package sortingoutpost

import "testing"

func TestUniqueDigitProducts(t *testing.T) {
	tests := []struct {
		a []int
		x int
	}{
		{
			[]int{2, 8, 121, 42, 222, 23},
			3,
		},
		{
			[]int{239},
			1,
		},
		{
			[]int{100, 101, 111},
			2,
		},
		{
			[]int{100, 23, 42, 239, 22339, 9999999, 456, 78, 228, 1488},
			10,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			10,
		},
	}

	for i, test := range tests {
		actual := uniqueDigitProducts(test.a)

		if actual != test.x {
			t.Errorf("%d - Got %d; Wanted %d", i+1, actual, test.x)
		}
	}
}
