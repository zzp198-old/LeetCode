package problems

import "strings"

func decodeMessage(key string, message string) string {
	hash := map[rune]rune{}
	cur := 'a'
	for _, char := range key {
		if char >= 'a' && char <= 'z' {
			if _, has := hash[char]; !has {
				hash[char] = cur
				cur++
			}
		}
	}
	b := strings.Builder{}
	for _, char := range message {
		if char >= 'a' && char <= 'z' {
			b.WriteRune(hash[char])
		} else {
			b.WriteRune(char)
		}
	}
	return b.String()
}
