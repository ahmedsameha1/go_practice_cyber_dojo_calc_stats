package calcstats

import (
	"errors"
	"math"
)

var emptysliceerror error = errors.New("This is an empty slice!")

func Min(list []int) (int, error) {
	if len(list) == 0 {
		return math.MinInt, emptysliceerror
	}
	return list[0], nil
}