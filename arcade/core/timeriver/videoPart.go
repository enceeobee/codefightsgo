package timeriver

import (
	"fmt"
	"math/big"
)

func videoPart(part string, total string) []int {
	var partHr, partMin, partSec, totalHr, totalMin, totalSec int

	fmt.Sscanf(part, "%2d:%2d:%2d", &partHr, &partMin, &partSec)
	partS := (partHr * 3600) + (partMin * 60) + partSec

	fmt.Sscanf(total, "%2d:%2d:%2d", &totalHr, &totalMin, &totalSec)
	totalS := (totalHr * 3600) + (totalMin * 60) + totalSec

	rat := big.NewRat(int64(partS), int64(totalS))

	return []int{int(rat.Num().Int64()), int(rat.Denom().Int64())}
}
