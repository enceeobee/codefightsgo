package springofintegration

func stringsCrossover(inputArray []string, result string) int {
	var works bool
	var s1, s2 string
	r := 0

	// BRUTE
	for i := 0; i < len(inputArray)-1; i++ {
		for j := i + 1; j < len(inputArray); j++ {
			// Check if these two strings can combine to make result
			// Loop over each char, check that one of them is the letter at that index in result
			s1 = inputArray[i]
			s2 = inputArray[j]
			works = false

			for k := 0; k < len(s1); k++ {
				works = s1[k] == result[k] || s2[k] == result[k]
				if !works {
					break
				}
			}

			if works {
				r++
			}
		}
	}

	return r
}

// func stringsCrossover(inputArray []string, result string) int {
// 	pairs := []int{}
// 	cMap := make([][]int, len(result))

// 	// TODO - Handle the case when len(inputArray[0]) == 1

// 	/**
// 	{
// 		[]string{
// 			"aa",
// 			"ab",
// 			"ba"
// 		},
// 		"bb",
// 		1,
// 	},

// 	// Map which strings contain the proper chars
// 	[
// 		[2],
// 		[1],
// 	]

// 	{
// 		[]string{
// 			"abc",
// 			"aaa",
// 			"aba",
// 			"bab"},
// 		"bbb",
// 		2,
// 	},
// 	cMap = [
// 		[3],
// 		[0,2],
// 		[3]
// 	]

// 	// Then (?) filter out redundant matches

// 	// Return len(pairs)
// 	*/

// 	for i := 0; i < len(inputArray[0]); i++ {
// 		for j := 0; j < len(inputArray); j++ {
// 			char := inputArray[j][i]
// 			if char == result[i] {
// 				// add it to the cmap
// 				cMap[i] = append(cMap[i], j)
// 			}
// 		}
// 	}

// 	fmt.Println(cMap)

// 	// Make matches
// 	/*
// 		matches = {
// 			3: [0, 2]
// 		}
// 	*/
// 	var matches map[int][]int

// 	fmt.Println(matches)

// 	for _, inputs := range cMap {
// 		if _, ok :=
// 	}

// 	return len(pairs)
// }
