package springofintegration

import "math"

func pairOfShoes(shoes [][]int) bool {
	if math.Mod(float64(len(shoes)), 2) != 0 {
		return false
	}

	var size int
	var typ int
	myShoes := map[int]map[int]int{}

	for _, shoe := range shoes {
		typ = shoe[0]
		size = shoe[1]

		if myShoes[size] == nil {
			myShoes[size] = map[int]int{}
		}

		myShoes[size][typ]++

		if myShoes[size][0] > 0 && myShoes[size][1] > 0 {
			myShoes[size][0]--
			myShoes[size][1]--

			if myShoes[size][0] == 0 && myShoes[size][1] == 0 {
				delete(myShoes, size)
			}
		}
	}

	return len(myShoes) == 0
}

/*
func pairOfShoes(shoes [][]int) bool {
	if math.Mod(float64(len(shoes)), 2) != 0 {
		return false
	}

	// TODO - Keep singles as an int

	var size int
	var typ int
	singles := map[int]bool{}
	myShoes := map[int]map[int]bool{}

	for _, shoe := range shoes {
		typ = shoe[0]
		size = shoe[1]

		if myShoes[size] == nil {
			singles[size] = true
			myShoes[size] = map[int]bool{} // TODO - make this map[int]int{} to keep track of how many shoes of each type we have
		}

		myShoes[size][typ] = true
		// TODO -
		// if myShoes[size][0] > 0 && myShoes[size][1] > 0 {
		if myShoes[size][0] && myShoes[size][1] {
			delete(myShoes, size)
			delete(singles, size)
		}
	}

	// fmt.Println("myShoes", myShoes)
	// fmt.Println("singles", singles)

	return len(singles) == 0
}
*/

// func pairOfShoes(shoes [][]int) bool {
// 	var size int
// 	var typ int
// 	singles := map[int]bool{}
// 	myShoes := map[int]map[int]bool{}

// 	if math.Mod(float64(len(shoes)), 2) != 0 {
// 		return false
// 	}

// 	for _, shoe := range shoes {
// 		typ = shoe[0]
// 		size = shoe[1]

// 		// if myShoes[size] == nil {
// 		// 	singles[size] = true
// 		// 	myShoes[size] = map[int]bool{
// 		// 		0: typ == 0,
// 		// 		1: typ == 1,
// 		// 	}
// 		// } else {
// 		// 	myShoes[size][typ] = true
// 		// 	if myShoes[size][0] && myShoes[size][1] {
// 		// 		delete(myShoes, size)
// 		// 		delete(singles, size)
// 		// 	} else {
// 		// 		singles[size] = true
// 		// 	}
// 		// }

// 		if myShoes[size] == nil {
// 			singles[size] = true
// 			myShoes[size] = map[int]bool{}
// 		}

// 		myShoes[size][typ] = true
// 		if myShoes[size][0] && myShoes[size][1] {
// 			delete(myShoes, size)
// 			delete(singles, size)
// 			// } else {
// 			// 	singles[size] = true
// 		}

// 		// if myShoes[size][0] && myShoes[size][1] {
// 		// 	delete(singles, size)
// 		// 	delete(myShoes, size)
// 		// }
// 	}

// 	fmt.Println("myShoes", myShoes)
// 	fmt.Println("singles", singles)

// 	return len(singles) == 0
// }

// package codefightsgo

// import (
// 	"math"
// )

// func pairOfShoes(shoes [][]int) bool {
// 	var size int
// 	var typ int
// 	type feet struct {
// 		left  bool
// 		right bool
// 	}
// 	singles := map[int]bool{}
// 	myShoes := map[int]feet{}

// 	if math.Mod(float64(len(shoes)), 2) != 0 {
// 		return false
// 	}

// 	for _, shoe := range shoes {
// 		typ = shoe[0]
// 		size = shoe[1]

// 		if !myShoes[size].left && !myShoes[size].right {
// 			singles[size] = true
// 			myShoes[size] = feet{typ == 0, typ == 1}
// 		} else {
// 			// If we already have a pair in this size, reset
// 			if myShoes[size].left && myShoes[size].right {
// 				delete(singles, size)
// 				myShoes[size] = feet{typ == 0, typ == 1}
// 			}

// 			singles[size] = true
// 			if typ == 0 {
// 				myShoes[size] = feet{true, myShoes[size].right}
// 			} else {
// 				myShoes[size] = feet{myShoes[size].left, true}
// 			}

// 			if myShoes[size].left && myShoes[size].right {
// 				delete(singles, size)
// 			}
// 		}
// 	}

// 	return len(singles) == 0
// }
