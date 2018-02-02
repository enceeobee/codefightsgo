package codefightsgo

import "strconv"

func fileNaming(names []string) []string {
	var rename string
	memo := make(map[string]bool)
	renames := []string{}

	for _, name := range names {
		if _, ok := memo[name]; !ok {
			rename = name
		} else {
			for i := 1; i > 0; i++ {
				rename = name + "(" + strconv.Itoa(i) + ")"
				if _, ok := memo[rename]; !ok {
					break
				}
			}
		}

		memo[rename] = true
		renames = append(renames, rename)
	}

	return renames
}
