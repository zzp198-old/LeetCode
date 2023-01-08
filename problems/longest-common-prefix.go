package problems

import "strings"

func longestCommonPrefix(strs []string) string {
	bd := strings.Builder{}
	for i := 0; i < len(strs[0]); i++ {
		prx := strs[0][i]
		for _, str := range strs {
			if i >= len(str) || str[i] != prx {
				return bd.String()
			}
		}
		bd.WriteRune(rune(prx))
	}
	return bd.String()
}
