package arrays

import "slices"

func AreOrderHistoriesEqual(history1, history2 [][]string) bool {

	if history1 == nil || history2 == nil {
		return false
	}

	return slices.EqualFunc(history1, history2, func(h1, h2 []string) bool {
		return slices.Equal(h1, h2)
	})
}
