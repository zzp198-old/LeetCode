package problems

func greatestLetter(s string) string {
	hash := map[rune]bool{}
	for _, c := range s {
		hash[c] = true
	}
	for i := 'Z'; i >= 'A'; i-- {
		if hash[i] && hash[i+32] {
			return string(i)
		}
	}
	return ""
}
