package springofintegration

import "testing"

func TestRunnersMeeting(t *testing.T) {
	var actual int
	tests := []struct {
		pos   []int
		speed []int
		x     int
	}{
		{[]int{1, 4, 2}, []int{27, 18, 24}, 3},
		{[]int{1, 4, 2}, []int{5, 6, 2}, 2},
		{[]int{1, 2, 3}, []int{1, 1, 1}, -1},
		{[]int{1, 1000}, []int{23, 22}, 2},
	}

	for _, test := range tests {
		actual = runnersMeetings(test.pos, test.speed)
		if actual != test.x {
			t.Errorf("runnersMeetings(%v, %d) = %d; expected %d", test.pos, test.speed, actual, test.x)
		}
	}
}
