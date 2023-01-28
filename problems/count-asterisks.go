package problems

func countAsterisks(s string) int {
	answer := 0
	in := false
	for _, c := range s {
		if c == '|' {
			in = !in
		}
		if c == '*' && !in {
			answer++
		}
	}
	return answer
}
