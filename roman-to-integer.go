package LeetCode

func romanToInt(s string) (answer int) {
	for i := 0; i < len(s); i++ {
		if s[i] == 'M' {
			answer += 1000
		} else if s[i] == 'D' {
			answer += 500
		} else if s[i] == 'C' {
			if i+1 < len(s) {
				if s[i+1] == 'D' {
					answer += 400
					i++
					continue
				} else if s[i+1] == 'M' {
					answer += 900
					i++
					continue
				}
			}
			answer += 100
		} else if s[i] == 'L' {
			answer += 50
		} else if s[i] == 'X' {
			if i+1 < len(s) {
				if s[i+1] == 'L' {
					answer += 40
					i++
					continue
				} else if s[i+1] == 'C' {
					answer += 90
					i++
					continue
				}
			}
			answer += 10
		} else if s[i] == 'V' {
			answer += 5
		} else if s[i] == 'I' {
			if i+1 < len(s) {
				if s[i+1] == 'V' {
					answer += 4
					i++
					continue
				} else if s[i+1] == 'X' {
					answer += 9
					i++
					continue
				}
			}
			answer += 1
		}
	}
	return
}
