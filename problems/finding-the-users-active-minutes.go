package problems

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	mp := map[int]map[int]bool{}
	for _, log := range logs {
		if mp[log[0]] == nil {
			mp[log[0]] = map[int]bool{}
		}
		mp[log[0]][log[1]] = true
	}
	answer := make([]int, k+1)
	for _, m := range mp {
		answer[len(m)]++
	}
	return answer[1:]
}
