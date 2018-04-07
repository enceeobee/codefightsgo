package timeriver

import (
	"reflect"
	"testing"
)

func TestVideoPart(t *testing.T) {
	tests := []struct {
		p string
		t string
		x []int
	}{
		{"02:20:00", "7:00:00", []int{1, 3}},
		{"25:26:20", "25:26:20", []int{1, 1}},
		{"00:02:20", "00:10:00", []int{7, 30}},
	}

	var actual = make([]int, 2)
	for _, test := range tests {
		actual = videoPart(test.p, test.t)
		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("videoPart(%s, %s) = %v; wanted %v", test.p, test.t, actual, test.x)
		}
	}
}
