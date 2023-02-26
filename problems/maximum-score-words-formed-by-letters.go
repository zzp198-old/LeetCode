package problems

func maxScoreWords(words []string, letters []byte, score []int) (ans int) {
	left := [26]int{}
	for _, c := range letters {
		left[c-'a']++
	}

	var dfs func(int, int)
	dfs = func(i, total int) {
		if i < 0 { // base case
			ans = max(ans, total)
			return
		}

		// 不选 words[i]
		dfs(i-1, total)

		// 选 words[i]
		for j, c := range words[i] {
			c -= 'a'
			if left[c] == 0 { // 剩余字母不足
				for _, c := range words[i][:j] { // 撤销
					left[c-'a']++
				}
				return
			}
			left[c]--         // 减少剩余字母
			total += score[c] // 累加得分
		}

		dfs(i-1, total)

		// 恢复现场
		for _, c := range words[i] {
			left[c-'a']++
		}
	}
	dfs(len(words)-1, 0)
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
