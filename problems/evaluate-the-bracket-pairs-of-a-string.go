package problems

import (
	"strings"
)

//import (
//	"regexp"
//	"strings"
//)
//
//func evaluate(s string, knowledge [][]string) string {
//	for _, know := range knowledge {
//		k := "(" + know[0] + ")"
//		v := know[1]
//		s = strings.ReplaceAll(s, k, v)
//	}
//	re, _ := regexp.Compile("\\((.*?)\\)")
//	return re.ReplaceAllString(s, "?")
//}

func evaluate(s string, knowledge [][]string) string {
	dict := make(map[string]string)
	for _, key := range knowledge {
		dict[key[0]] = key[1]
	}
	answer := strings.Builder{}
	start := -1
	for i, c := range s {
		if c == '(' {
			start = i
		} else if c == ')' {
			if t, ok := dict[s[start+1:i]]; ok {
				answer.WriteString(t)
			} else {
				answer.WriteByte('?')
			}
			start = -1
		} else if start < 0 {
			answer.WriteRune(c)
		}
	}
	return answer.String()
}
