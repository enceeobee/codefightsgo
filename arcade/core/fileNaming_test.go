package codefightsgo

import "testing"
import "reflect"

func TestFileNaming(t *testing.T) {
	tests := []struct {
		names []string
		x     []string
	}{
		{[]string{
			"doc",
			"doc",
			"image",
			"doc(1)",
			"doc",
		}, []string{
			"doc",
			"doc(1)",
			"image",
			"doc(1)(1)",
			"doc(2)",
		}},

		{[]string{
			"a(1)",
			"a(6)",
			"a",
			"a",
			"a",
			"a",
			"a",
			"a",
			"a",
			"a",
			"a",
			"a",
		}, []string{
			"a(1)",
			"a(6)",
			"a",
			"a(2)",
			"a(3)",
			"a(4)",
			"a(5)",
			"a(7)",
			"a(8)",
			"a(9)",
			"a(10)",
			"a(11)",
		}},

		{[]string{
			"dd",
			"dd(1)",
			"dd(2)",
			"dd",
			"dd(1)",
			"dd(1)(2)",
			"dd(1)(1)",
			"dd",
			"dd(1)",
		}, []string{
			"dd",
			"dd(1)",
			"dd(2)",
			"dd(3)",
			"dd(1)(1)",
			"dd(1)(2)",
			"dd(1)(1)(1)",
			"dd(4)",
			"dd(1)(3)",
		}},
	}

	for _, test := range tests {
		actual := fileNaming(test.names)
		if !reflect.DeepEqual(actual, test.x) {
			t.Errorf("fileNaming(%q) = %q; wanted %q", test.names, actual, test.x)
		}
	}
}
