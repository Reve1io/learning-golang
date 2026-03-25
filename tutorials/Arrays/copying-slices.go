package arrays

import (
	"fmt"
	"slices"
)

func IntsCopy(src []int, maxLen int) []int {

	if maxLen < 0 || maxLen == 0 {
		return []int{}
	} else if maxLen > len(src) {
		copy := slices.Clone(src)
		return copy
	}

	newSrc := make([]int, maxLen)
	fmt.Println(len(newSrc))

	copy(newSrc, src)

	fmt.Println(len(newSrc))

	return newSrc
}
