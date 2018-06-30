package timeriver

import "testing"

func TestNewYearCelebrations(t *testing.T) {
	var actual int
	tests := []struct {
		takeOffTime string
		minutes     []int
		x           int
	}{
		{"23:35", []int{60, 90, 140}, 3},
		{"00:00", []int{60, 120, 180, 250}, 4},
		{"13:36", []int{23, 33, 45, 56, 66, 88}, 1},
		{"22:50", []int{}, 1},
		{"20:18", []int{222, 342}, 3},
		{"12:05", []int{1, 109, 113, 344, 345, 478, 545, 565, 809, 814, 838, 883, 971, 1007, 1029, 1053, 1133, 1271, 1314, 1500}, 1},
		{"17:10", []int{150, 160, 293, 589, 614, 716, 760, 776, 781, 911, 1040, 1438}, 3},
		{"18:15", []int{117, 241, 364, 375, 545, 1317}, 1},
		{"19:44", []int{545, 1320}, 1},
		{"21:13", []int{52, 257, 323, 515, 579, 600, 703, 707, 1127, 1298}, 3},
	}

	for i, test := range tests {
		actual = newYearCelebrations(test.takeOffTime, test.minutes)

		if actual != test.x {
			t.Errorf("Test %d = %d; expected %d", i+1, actual, test.x)
		}
	}
}
