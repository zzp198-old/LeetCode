package problems

import (
	"math"
)

func rearrangeCharacters(s string, target string) int {
	dict1 := make(map[rune]int)
	for _, c := range s {
		dict1[c]++
	}
	dict2 := make(map[rune]int)
	for _, c := range target {
		dict2[c]++
	}
	answer := math.MaxInt
	for k, x := range dict2 {
		if y, ok := dict1[k]; ok {
			answer = int(math.Min(float64(answer), float64(y/x)))
		} else {
			return 0
		}
	}
	return answer
}
