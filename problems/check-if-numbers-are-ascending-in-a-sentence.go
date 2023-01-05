package problems

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	pre := 0
	for _, key := range strings.Split(s, " ") {
		if key[0] <= '9' {
			v, _ := strconv.Atoi(key)
			if pre >= v {
				return false
			}
			pre = v
		}
	}
	return true
}
