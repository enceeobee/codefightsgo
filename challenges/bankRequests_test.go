package challenges

import (
	"reflect"
	"testing"
)

func TestBankRequests(t *testing.T) {
	tests := []struct {
		accounts []int
		requests []string
		x        []int
	}{
		{
			[]int{10, 100, 20, 50, 30},
			[]string{"withdraw 2 10",
				"transfer 5 1 20",
				"deposit 5 20",
				"transfer 3 4 15"},
			[]int{30, 90, 5, 65, 30},
		},
	}

	for i, test := range tests {

		actual := bankRequests(test.accounts, test.requests)

		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("Failed test %d; got %v; wanted %v", i+1, actual, test.x)
		}
	}
}
