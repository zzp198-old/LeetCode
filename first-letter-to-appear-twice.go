package LeetCode

import "go/types"

func repeatedCharacter(s string) (answer byte) {
	hash := make(map[rune]interface{})
	for _, b := range s {
		if _, ok := hash[b]; ok {
			answer = byte(b)
			break
		} else {
			hash[b] = types.Interface{}
		}
	}
	return
}
