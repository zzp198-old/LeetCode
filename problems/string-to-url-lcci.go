package problems

import "strings"

func replaceSpaces(S string, length int) string {
	S = strings.ReplaceAll(S[:length], " ", "%20")
	return S
}
