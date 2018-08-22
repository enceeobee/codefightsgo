package chesstavern

import (
	"reflect"
	"testing"
)

func TestAmazonCheckmate(t *testing.T) {
	tests := []struct {
		king   string
		amazon string
		x      []int
	}{
		{"d3", "e4", []int{5, 21, 0, 29}},
		{"a1", "g5", []int{0, 29, 1, 29}},
		{"a3", "e4", []int{1, 32, 1, 23}},
		{"f3", "f2", []int{6, 11, 0, 38}},
		{"b7", "a8", []int{0, 10, 0, 45}},
		{"f7", "d3", []int{4, 28, 1, 21}},
		{"g2", "c3", []int{9, 21, 0, 24}},
		{"f3", "c1", []int{4, 18, 0, 32}},
		{"d4", "h8", []int{0, 18, 0, 36}},
	}

	var actual []int
	for i, test := range tests {
		actual = amazonCheckmate(test.king, test.amazon)

		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("Test %d failed\ngot\n%v\nwanted\n%v", i+1, actual, test.x)
		}
	}
}

func TestCanCaptureSameCol(t *testing.T) {
	tests := []struct {
		amazon    acPiece
		king      acPiece
		blackKing acPiece
		x         bool
	}{
		// Above
		{
			acPiece{4, 4},
			acPiece{0, 4},
			acPiece{1, 4},
			true,
		},
		// Above, but blocked by king
		{
			acPiece{4, 4},
			acPiece{1, 4},
			acPiece{0, 4},
			false,
		},
		// Below
		{
			acPiece{4, 4},
			acPiece{7, 4},
			acPiece{6, 4},
			true,
		},
		// Below, but blocked by king
		{
			acPiece{4, 4},
			acPiece{5, 4},
			acPiece{6, 4},
			false,
		},
	}

	var actual bool
	for _, test := range tests {

		actual = canCaptureSameCol(test.amazon, test.king, test.blackKing)

		if actual != test.x {
			t.Errorf("canCaptureSameCol(%v, %v, %v) = %t; expected %t", test.amazon, test.king, test.blackKing, actual, test.x)
		}
	}
}

func TestCanCaptureSameRow(t *testing.T) {
	tests := []struct {
		amazon    acPiece
		king      acPiece
		blackKing acPiece
		x         bool
	}{
		// Right
		{
			acPiece{4, 4},
			acPiece{4, 1},
			acPiece{4, 6},
			true,
		},
		// Right, but blocked by king
		{
			acPiece{4, 4},
			acPiece{4, 6},
			acPiece{4, 7},
			false,
		},
		// Left
		{
			acPiece{4, 4},
			acPiece{4, 5},
			acPiece{4, 0},
			true,
		},
		// Left, but blocked by king
		{
			acPiece{4, 4},
			acPiece{4, 2},
			acPiece{4, 1},
			false,
		},
	}

	var actual bool
	for _, test := range tests {
		actual = canCaptureSameRow(test.amazon, test.king, test.blackKing)

		if actual != test.x {
			t.Errorf("canCaptureSameRow(%v, %v, %v) = %t; expected %t", test.amazon, test.king, test.blackKing, actual, test.x)
		}
	}
}

func TestCanCaptureDiagonal(t *testing.T) {
	tests := []struct {
		amazon    acPiece
		king      acPiece
		blackKing acPiece
		x         bool
	}{
		{
			acPiece{row: 0, col: 0},
			acPiece{row: 1, col: 1},
			acPiece{2, 2},
			false,
		},
		{
			acPiece{row: 4, col: 4},
			acPiece{row: 3, col: 3},
			acPiece{1, 1},
			false,
		},
		{
			acPiece{row: 4, col: 4},
			acPiece{row: 5, col: 5},
			acPiece{1, 1},
			true,
		},
		{
			acPiece{4, 4},
			acPiece{5, 3},
			acPiece{0, 0},
			true,
		},
		{
			acPiece{4, 4},
			acPiece{5, 3},
			acPiece{1, 7},
			true,
		},
	}

	var actual bool
	for _, test := range tests {
		actual = canCaptureDiagonal(test.amazon, test.king, test.blackKing)
		if actual != test.x {
			t.Errorf("canCaptureDiagonal(%v, %v, %v, %d) = %t; expected %t", test.amazon, test.king, test.blackKing, actual, test.x)
		}
	}
}
