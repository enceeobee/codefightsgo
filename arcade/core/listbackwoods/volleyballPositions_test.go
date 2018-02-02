package listbackwoods

import "testing"
import "reflect"

func TestVolleyballPositions(t *testing.T) {
	tests := []struct {
		formation [][]string
		k         int
		x         [][]string
	}{
		{
			[][]string{
				[]string{"empty", "Player5", "empty"},
				[]string{"Player4", "empty", "Player2"},
				[]string{"empty", "Player3", "empty"},
				[]string{"Player6", "empty", "Player1"},
			},
			2,
			[][]string{
				[]string{"empty", "Player1", "empty"},
				[]string{"Player2", "empty", "Player3"},
				[]string{"empty", "Player4", "empty"},
				[]string{"Player5", "empty", "Player6"},
			},
		},
	}

	for _, test := range tests {
		actual := volleyballPositions(test.formation, test.k)

		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("garrrr")
		}
	}
}
