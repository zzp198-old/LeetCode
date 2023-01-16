package problems

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	word1 := strings.Split(sentence1, " ")
	word2 := strings.Split(sentence2, " ")
	i, n := 0, len(word1)
	j, m := 0, len(word2)
	for i < n && i < m && word1[i] == word2[i] {
		i++
	}
	for j < n-i && j < m-i && word1[n-j-1] == word2[m-j-1] {
		j++
	}
	return i+j == min(n, m)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
