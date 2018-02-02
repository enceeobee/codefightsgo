package springofintegration

import "testing"

func TestPairOfShoes(t *testing.T) {
	tests := []struct {
		shoes    [][]int
		expected bool
	}{
		{[][]int{
			{0, 21},
			{1, 23},
			{1, 21},
			{0, 23}}, true},
		{[][]int{
			{1, 21},
			{1, 21},
			{1, 21}}, false},
		{[][]int{
			{1, 21},
			{1, 21},
			{1, 21},
			{1, 21},
			{0, 21}}, false},
		{[][]int{
			{1, 21},
			{1, 21},
			{1, 21},
			{1, 21}}, false},
		{[][]int{
			{0, 21},
			{0, 21},
			{1, 21},
			{1, 21}}, true},
		{[][]int{
			{0, 21},
			{1, 21},
			{0, 21},
			{1, 21}}, true},
		{[][]int{
			{0, 21},
			{1, 21},
			{1, 21}}, false},
		{[][]int{
			{0, 21},
			{1, 21},
			{1, 21},
			{0, 21}}, true},
		{[][]int{
			{0, 21},
			{1, 21},
			{1, 21},
			{0, 22}}, false},
		{[][]int{
			{0, 21},
			{1, 23},
			{1, 21},
			{1, 23}}, false},
		{[][]int{
			{0, 23},
			{1, 21},
			{1, 23},
			{0, 21},
			{1, 22},
			{0, 22}}, true},
		{[][]int{
			{0, 23},
			{1, 21},
			{1, 23},
			{0, 21}}, true},
		{[][]int{
			{0, 23},
			{1, 21},
			{1, 23},
			{0, 21}}, true},
		{[][]int{
			{0, 23}}, false},
		{[][]int{
			{0, 23},
			{1, 23}}, true},
		{[][]int{
			{0, 23},
			{1, 22}}, false},
		{[][]int{
			{0, 23},
			{1, 22},
			{1, 22}}, false},
	}

	actual := false

	for i, test := range tests {
		actual = pairOfShoes(test.shoes)
		if actual != test.expected {
			t.Errorf("Error in test %d; wanted %t, got %t", i+1, test.expected, actual)
		}
	}
}
