package arrays

import (
	"slices"
)

func AreOrderHistoriesEqual(history1, history2 [][]string) bool {
	var result bool

	if history1 == nil || history2 == nil {
		return false
	}

	history11, history12 := history1[0], history1[1]
	history21, history22 := history2[0], history2[1]

	equalFirstLayer := slices.EqualFunc(history11, history21, func(h11, h21 string) bool {
		if h11 == h21 {
			return true
		}
		return false
	})

	equalSecondLayer := slices.EqualFunc(history12, history22, func(h12, h22 string) bool {
		if h12 == h22 {
			return true
		}
		return false
	})

	if equalFirstLayer && equalSecondLayer == true {
		result = true
	} else {
		result = false
	}

	return result
}
