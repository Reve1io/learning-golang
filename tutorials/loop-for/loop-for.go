package loopfor

import "strings"

func RepeatWord(word string, times int) string {
	var storageWord string

	if times != 1 {
		for i := 1; i <= times; i++ {
			storageWord = strings.Repeat(word, i)
		}
		return storageWord
	} else {
		return word
	}
}
