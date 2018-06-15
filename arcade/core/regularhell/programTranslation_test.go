package regularhell

import "testing"

func TestProgramTranslation(t *testing.T) {
	tests := []struct {
		s string
		a []string
		x string
	}{
		{"function add($n, m) {\t  return n + $m;\t}", []string{"n", "m"}, "function add($n, $m) {\t  return $n + $m;\t}"},
		{"function findSum(a, $cnt) {\t  var a0 = $a;\t  for (var _cnt = 0, _cnt < cnt; _cnt++)\t    a0 += _cnt;\t  return a0;\t}", []string{"a", "cnt"}, "function findSum($a, $cnt) {\t  var a0 = $a;\t  for (var _cnt = 0, _cnt < $cnt; _cnt++)\t    a0 += _cnt;\t  return a0;\t}"},
		{"function doNothing($uselessVariable) {\t  return $uselessVariable;\t}", []string{"uselessVariable"}, "function doNothing($uselessVariable) {\t  return $uselessVariable;\t}"},
		{"function addToVariable(variable) {\t  variable_which_should_be_increased_by_the_variable = 14;\t  variable_which_should_be_increased_by_the_variable += variable;\t  return variable_which_should_be_increased_by_the_variable;\t}", []string{"variable"}, "function addToVariable($variable) {\t  variable_which_should_be_increased_by_the_variable = 14;\t  variable_which_should_be_increased_by_the_variable += $variable;\t  return variable_which_should_be_increased_by_the_variable;\t}"},
		{"function replaceThemAll(rep, laceT, hemAll, ornot) {\t  var tmp = rep;\t  rep = laceT;\t  laceT = hemAll;\t  hemAll = tmp;\t  return [rep, laceT, hemAll]\t}", []string{"rep", "laceT", "hemAll"}, "function replaceThemAll($rep, $laceT, $hemAll, ornot) {\t  var tmp = $rep;\t  $rep = $laceT;\t  $laceT = $hemAll;\t  $hemAll = tmp;\t  return [$rep, $laceT, $hemAll]\t}"},
		{"function returnSecond(fu_,_re5,NOO) {\t  return _re5;\t}", []string{"fu_", "_re5", "NOO"}, "function returnSecond($fu_,$_re5,$NOO) {\t  return $_re5;\t}"},
		{"function getLength(k, m) {\t  return m.length;\t}", []string{"m"}, "function getLength(k, $m) {\t  return $m.length;\t}"},
	}

	var actual string

	for _, test := range tests {
		actual = programTranslation(test.s, test.a)

		if actual != test.x {
			t.Errorf("programTranslation(%v, %v) = %v; expected %v", test.s, test.a, actual, test.x)
		}
	}
}
