package challenges

import (
	"strconv"
	"strings"
)

func bankRequests(accounts []int, requests []string) []int {
	var args []string
	var reqType string

	for i, request := range requests {
		args = strings.Split(request, " ")
		reqType = args[0]

		if reqType == "withdraw" {
			withdrawAccount, err := strconv.Atoi(args[1])

			if err != nil {
				return []int{-(i + 1)}
			}
			withdrawAccount--

			if withdrawAccount < 0 || withdrawAccount > len(accounts)-1 {
				return []int{-(i + 1)}
			}

			withdrawAmt, _ := strconv.Atoi(args[2])

			accounts[withdrawAccount] -= withdrawAmt

			if accounts[withdrawAccount] < 0 {
				return []int{-(i + 1)}
			}
		} else if reqType == "deposit" {
			depositAccount, err := strconv.Atoi(args[1])

			if err != nil {
				return []int{-(i + 1)}
			}

			depositAccount--

			if depositAccount < 0 || depositAccount > len(accounts)-1 {
				return []int{-(i + 1)}
			}

			depositAmt, _ := strconv.Atoi(args[2])
			accounts[depositAccount] += depositAmt
		} else {
			withdrawAccount, err := strconv.Atoi(args[1])

			if err != nil {
				return []int{-(i + 1)}
			}

			withdrawAccount--

			if withdrawAccount < 0 || withdrawAccount > len(accounts)-1 {
				return []int{-(i + 1)}
			}

			withdrawAmt, _ := strconv.Atoi(args[3])

			depositAccount, err := strconv.Atoi(args[2])

			if err != nil {
				return []int{-(i + 1)}
			}

			depositAccount--

			if depositAccount < 0 || depositAccount > len(accounts)-1 {
				return []int{-(i + 1)}
			}

			accounts[withdrawAccount] -= withdrawAmt

			if accounts[withdrawAccount] < 0 {
				return []int{-(i + 1)}
			}
			accounts[depositAccount] += withdrawAmt
		}
	}

	return accounts
}
