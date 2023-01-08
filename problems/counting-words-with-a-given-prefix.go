package problems

import "strings"

func prefixCount(words []string, pref string) (answer int) {
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			answer++
		}
	}
	return
}
