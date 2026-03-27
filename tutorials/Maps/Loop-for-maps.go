package maps

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
)

func CountWords(text string) map[string]int {

	if text == "" {
		return make(map[string]int)
	}

	text = strings.ToLower(text)

	var result strings.Builder
	for _, r := range text {
		if !unicode.IsPunct(r) {
			result.WriteRune(r)
		}
	}

	text = result.String()
	words := strings.Fields(text)
	words = slices.Compact(words)

	repeat := make(map[string]int)

	for _, word := range words {
		count := strings.Count(text, word)
		repeat[word] = count
	}

	fmt.Println(words)
	return repeat
}
