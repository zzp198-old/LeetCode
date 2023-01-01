package LeetCode

func repeatedCharacter(s string) (answer byte) {
	hash := make(map[byte]interface{})
	for _, b := range s {
		if _, ok := hash[byte(b)]; ok {
			answer = byte(b)
			break
		} else {
			hash[byte(b)] = 1
		}
	}
	return
}
